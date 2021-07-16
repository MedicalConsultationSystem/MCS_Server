package request

type AddPrescription struct {
	OrgId            int    `json:"org_id"`
	UserId           string `json:"user_id"`
	ConsultId        int    `json:"consult_id"`
	PrescriptionType string `json:"prescription_type"`
	DoctorId         string `json:"doctor_id"`
	DoctorName       string `json:"doctor_name"`
}
