package request

type AddOrg struct {
	OrgName string `json:"org_name"`
}

type FindOrg struct {
	OrgName string `json:"org_name"`
}