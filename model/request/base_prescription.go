package request

type AddPrescription struct {
	OrgId            int    `json:"org_id"`
	UserId           string `json:"user_id"`
	ConsultId        int    `json:"consult_id"`
	PrescriptionType string `json:"prescription_type"`
	DoctorId         string `json:"doctor_id"`
	DoctorName       string `json:"doctor_name"`
}

type PrescriptionDrug struct {
	OrgId          int     `json:"org_id"`
	PrescriptionId int     `json:"prescription_id"`
	DrugId         int     `json:"drug_id"`
	DrugName       string  `json:"drug_name"`
	Specification  string  `json:"specification"`
	Dose           float64 `json:"dose"`
	DoseUnit       string  `json:"dose_unit"`
	FrequencyCode  int     `json:"frequency_code"`
	FrequencyName  string  `json:"frequency_name"`
	UsageCode      int     `json:"usage_code"`
	UsageName      string  `json:"usage_name"`
	TakeDays       int     `json:"take_days"`
	Quantity       float64 `json:"quantity"`
	Price          float64 `json:"price"`
	PackUnit       string  `json:"pack_unit"`
	GroupNumber    int     `json:"group_number"`
	SortNumber     int     `json:"sort_number"`
	Remark         string  `json:"remark"`
}

//type Prescription struct {
//	PrescriptionId int                `json:"prescription_id"`
//	Drugs          []PrescriptionDrug `json:"drugs"`
//}

type SubmitPrescription struct {
	PrescriptionIds []int `json:"prescription_ids"`
}

type DeletePrescription struct {
	PrescriptionDrugId int `json:"prescription_drug_id"`
}

type ListPrescription struct {
	ConsultId int `json:"consult_id"`
}
