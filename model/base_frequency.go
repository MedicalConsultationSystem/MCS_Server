package model

type BaseFrequency struct {
	ItemCode int `json:"item_code"`
	ItemName string `json:"item_name"`
	ItemNameAbbr string `json:"item_name_abbr"`
}

func (m *BaseFrequency) TableName() string {
	return "base_dic_drug_frequency"
}