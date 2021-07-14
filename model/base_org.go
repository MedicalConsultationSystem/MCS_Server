package model

type BaseOrg struct {
	OrgId   int    `json:"org_id"`
	OrgName string `json:"org_name"`
}
func (m *BaseOrg) TableName() string {
	return "base_org"
}