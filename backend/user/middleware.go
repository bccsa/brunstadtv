package user

import (
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/gin-gonic/gin"
)

// All well-known roles as used in the DB
const (
	RolePublic       = "public"
	RoleRegistered   = "registered"
	RoleBCCMember    = "bcc-members"
	RoleNonBCCMember = "non-bcc-members"
)

// Various hardcoded keys
const (
	CtxUser          = "ctx-user"
	CacheRoles       = "roles"
	CtxLanguages     = "ctx-languages"
	CtxImpersonating = "ctx-impersonating"
	CtxProfiles      = "ctx-profiles"
	CtxProfile       = "ctx-profile"
)

// GetAcceptedLanguagesFromCtx as sent by the user
func GetAcceptedLanguagesFromCtx(ctx *gin.Context) []string {
	accLang := ctx.GetHeader("Accept-Language")
	return utils.ParseAcceptLanguage(accLang)
}

// AgeGroups contains the different age groups keyed by the minimum age.
var AgeGroups = map[int]string{
	65: "65+",
	51: "51 - 64",
	37: "37 - 50",
	26: "26 - 36",
	19: "19 - 25",
	13: "13 - 18",
	10: "10 - 12",
	0:  "0 - 9",
}

func ageFromBirthDate(birthDate string) int {
	date, err := time.Parse("2006-04-02", birthDate)
	if err != nil {
		return 0
	}
	now := time.Now()

	// years since birth
	years := now.Year() - date.Year()

	// if the user hasn't had their birthday yet this year, subtract a year
	if now.Month() < date.Month() || (now.Month() == date.Month() && now.Day() < date.Day()) {
		years--
	}

	return years
}

// NewUserMiddleware returns a gin middleware that ingests a populated User struct
// into the gin context
func NewUserMiddleware(queries *sqlc.Queries, remoteCache *remotecache.Client, ls *common.BatchLoaders, auth0Client *auth0.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		reqCtx, span := otel.Tracer("user/middleware").Start(ctx.Request.Context(), "run")
		defer span.End()

		var roles []string

		authed := ctx.GetBool(auth0.CtxAuthenticated)

		// This can't be on the user object because that is cached for too long
		ctx.Set(CtxLanguages, GetAcceptedLanguagesFromCtx(ctx))

		// If the user is anonymous we just create a simple object and bail
		if !authed {
			span.AddEvent("Anonymous")
			roles = append(roles, RolePublic)
			ctx.Set(CtxUser,
				&common.User{
					Roles:     roles,
					Anonymous: true,
					ActiveBCC: false,
				})
			return
		}

		userID := ctx.GetString(auth0.CtxUserID)

		if u, ok := userCache.Get(userID); ok {
			span.AddEvent("User From Cache")
			ctx.Set(CtxUser, u)
			return
		}

		getUserFromMembers := func(o *remotecache.Options) (*common.User, error) {
			o.SetTTL(time.Minute * 5)
			pid := ctx.GetString(auth0.CtxPersonID)
			intID, _ := strconv.ParseInt(pid, 10, 32)

			roles = append(roles, RoleRegistered)
			if ctx.GetBool(auth0.CtxIsBCCMember) {
				roles = append(roles, RoleBCCMember)
			} else {
				roles = append(roles, RoleNonBCCMember, RolePublic)
			}

			if pid == "" || pid == "0" {
				pid = ctx.GetString(auth0.CtxUserID)
			}

			u := &common.User{
				PersonID:  pid,
				Roles:     roles,
				Anonymous: false,
				ActiveBCC: ctx.GetBool(auth0.CtxIsBCCMember),
				AgeGroup:  "unknown",
				Gender:    "unknown",
			}

			saveUser := func() error {
				return queries.UpsertUser(ctx, sqlc.UpsertUserParams{
					ID:            u.PersonID,
					Roles:         u.Roles,
					DisplayName:   u.DisplayName,
					FirstName:     u.FirstName,
					ActiveBcc:     u.ActiveBCC,
					EmailVerified: u.EmailVerified,
					Email:         u.Email,
					AgeGroup:      u.AgeGroup,
					Age:           int32(u.Age),
					Gender:        u.Gender,
					ChurchIds: lo.Map(u.ChurchIDs, func(i int, _ int) int32 {
						return int32(i)
					}),
				})
			}

			//info, err := auth0Client.GetUserInfoForAuthHeader(ctx, ctx.GetHeader("Authorization"))
			//if err != nil {
			//	return nil, err
			//}
			//u.EmailVerified = info.EmailVerified
			u.EmailVerified = true

			if u.IsActiveBCC() {
				member, err := ls.MemberLoader.Get(ctx, int(intID))
				if err != nil {
					log.L.Info().Err(err).Msg("Failed to retrieve user from members.")
					span.AddEvent("User failed to load from members")

					dbUser, err := ls.UserLoader.Get(ctx, pid)
					if err != nil {
						log.L.Error().Err(err).Msg("Failed to retrieve user from database.")
					}
					if dbUser != nil {
						u = dbUser
					}

					userCache.Set(userID, u, cache.WithExpiration(1*time.Minute))
					ctx.Set(CtxUser, u)
					return u, nil
				}
				u.FirstName = member.FirstName
				switch member.Gender {
				case "Male":
					u.Gender = "male"
				case "Female":
					u.Gender = "female"
				default:
					u.Gender = "unknown"
				}
				u.Email = member.Email
				u.DisplayName = member.DisplayName

				u.Age = ageFromBirthDate(member.BirthDate)

				now := time.Now()
				affiliations := lo.Filter(member.Affiliations, func(i members.Affiliation, _ int) bool {
					return (i.ValidTo == nil || i.ValidTo.After(now)) && (i.ValidFrom == nil || i.ValidFrom.Before(now))
				})
				organizations, err := ls.OrganizationLoader.GetMany(ctx, lo.Map(affiliations, func(i members.Affiliation, _ int) uuid.UUID {
					return i.OrgUid
				}))
				if err != nil {
					return nil, err
				}
				for _, org := range organizations {
					if org != nil && org.Type == "Church" {
						u.ChurchIDs = append(u.ChurchIDs, org.OrgID)
					}
				}
			} else {
				info, err := auth0Client.GetUser(ctx, ctx.GetString(auth0.CtxUserID))
				if err != nil {
					return nil, err
				}
				u.PersonID = info.UserId
				u.Email = info.Email
				u.DisplayName = info.Nickname
				u.EmailVerified = info.EmailVerified

				if by, ok := info.UserMetadata["birth_year"]; ok {
					year, err := strconv.ParseInt(by, 10, 64)
					if err == nil {
						u.Age = time.Now().Year() - int(year)
					}
				}
				u.CompletedRegistration = info.CompletedRegistration()
			}

			if u.Email == "" {
				// Explicit values make it easier to see that it was intended when debugging
				u.Email = "<MISSING>"
			}

			if u.EmailVerified {
				userRoles, err := GetRolesForEmail(reqCtx, queries, u.Email)
				if err != nil {
					err = merry.Wrap(err)
					log.L.Warn().Err(err).Str("email", u.Email).Msg("Unable to get roles")
				} else {
					roles = append(roles, userRoles...)
				}
			}

			u.Roles = roles

			ageGroupMin := 0
			for minAge, group := range AgeGroups {
				// Note: Maps are not iterated in a sorted order, so we have to find the lowed applicable
				if u.Age >= minAge && minAge > ageGroupMin {
					u.AgeGroup = group
					ageGroupMin = minAge
				}
			}

			err := saveUser()

			if err != nil {
				log.L.Error().Err(err).Send()
			}

			return u, nil
		}

		lock, _ := userCacheLocks.LoadOrStore(userID, &sync.Mutex{})
		lock.Lock()
		defer func() {
			lock.Unlock()
			userCacheLocks.Delete(userID)
		}()
		if u, err := remotecache.GetOrCreate[*common.User](ctx, remoteCache, fmt.Sprintf("users:%s", userID), getUserFromMembers); err == nil {
			span.AddEvent("User loaded into cache")
			userCache.Set(userID, u, cache.WithExpiration(60*time.Second))
			ctx.Set(CtxUser, u)
			return
		} else {
			log.L.Error().Err(err).Send()
			ctx.Set(CtxUser, &common.User{
				Roles:     roles,
				Anonymous: true,
				ActiveBCC: false,
			})
		}
	}
}

// GetFromCtx gets the user stored in the context by the middleware
func GetFromCtx(ctx *gin.Context) *common.User {
	u, ok := ctx.Get(CtxUser)
	if !ok {
		return nil
	}

	return u.(*common.User)
}

// GetProfileFromCtx returns the current profile
func GetProfileFromCtx(ctx *gin.Context) *common.Profile {
	p, ok := ctx.Get(CtxProfile)
	if !ok {
		return nil
	}
	return p.(*common.Profile)
}

// GetProfilesFromCtx returns the current profile
func GetProfilesFromCtx(ctx *gin.Context) []common.Profile {
	p, ok := ctx.Get(CtxProfiles)
	if !ok {
		return nil
	}
	return p.([]common.Profile)
}

// GetLanguagesFromCtx as provided in the request
func GetLanguagesFromCtx(ctx *gin.Context) []string {
	l, ok := ctx.Get(CtxLanguages)
	if !ok {
		return []string{}
	}

	return l.([]string)
}
