package members

type result[t any] struct {
	Data t `json:"data"`
}

// Member is a member with related data
type Member struct {
	PersonID     int
	Age          int
	Email        string
	DisplayName  string
	Affiliations []Affiliation
}

// Affiliation is an affiliation to an entity
type Affiliation struct {
	Active    bool
	OrgID     int
	OrgType   string
	Type      string
	ValidFrom string
	ChurchID  int
}

// Organization
type Organization struct {
	OrgID int
	Name  string `json:"districtName"`
	Type  string
}
