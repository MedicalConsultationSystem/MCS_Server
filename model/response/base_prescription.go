package response

import (
	"MCS_Server/model"
	"time"
)

type Prescription struct {
	PrescriptionId   int                          `json:"prescription_id"`
	OrgId            int                          `json:"org_id"`
	UserId           string                       `json:"user_id"`
	ConsultId        int                          `json:"consult_id"`
	PrescriptionType string                       `json:"prescription_type"`
	DoctorId         string                       `json:"doctor_id"`
	DoctorName       string                       `json:"doctor_name"`
	CreateTime       time.Time                    `json:"create_time"`
	Drugs            []model.BasePrescriptionDrug `json:"drugs"`
}

type ListPrescription struct {
	Prescriptions []Prescription `json:"prescriptions"`
}
