package notifications

import (
	"context"

	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/targets"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

// ResolveTargets resolves targetIDs to device tokens
func (u *Utils) ResolveTargets(ctx context.Context, targetIDs []uuid.UUID) ([]common.Device, error) {
	log.L.Debug().Int("targetCount", len(targetIDs)).Msg("Resolving targets")
	targetRows, err := u.queries.GetTargets(ctx, targetIDs)
	if err != nil {
		return nil, err
	}

	var devices []common.Device
	for _, t := range targetRows {
		target := common.Target(t)
		switch t.Type {
		case "usergroups":
			ds, err := u.getTokensForTarget(ctx, target)
			if err != nil {
				return nil, err
			}
			log.L.Debug().Int("deviceCount", len(ds)).Msg("Resolved target, retrieved devices")
			devices = append(devices, ds...)
		}
	}
	return devices, nil
}

func (u *Utils) getTokensForTarget(ctx context.Context, target common.Target) ([]common.Device, error) {
	apps, err := u.queries.ListApplications(ctx)
	if err != nil {
		return nil, err
	}
	defaultApp, _ := lo.Find(apps, func(i common.Application) bool {
		return i.Default
	})

	profileIDs, err := targets.ResolveProfileIDs(ctx, u.queries, defaultApp.GroupID, target)
	if err != nil {
		return nil, err
	}

	devices, err := u.queries.GetDevices(ctx, profileIDs)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
