package model

type BaseUsage struct {
	ItemCode int `json:"item_code"`
	ItemName string `json:"item_name"`
}
func (m *BaseUsage) TableName() string {
	return "base_dic_drug_usage"
}