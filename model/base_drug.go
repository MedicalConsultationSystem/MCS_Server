package model

type BaseDrug struct {
	DrugId         int     `json:"drug_id"`
	DrugName       string  `json:"drug_name"`
	TradeName      string  `json:"trade_name"`
	PinyinCode     string  `json:"pinyin_code"`
	Specification  string  `json:"specification"`
	PackUnit       string  `json:"pack_unit"`
	Price          float64 `json:"price"`
	Dose           float64 `json:"dose"`
	DoseUnit       string  `json:"dose_unit"`
	FactoryName    string  `json:"factory_name"`
	ApprovalNumber string  `json:"approval_number"`
}

func (m *BaseDrug) TableName() string {
	return "base_drug"
}
