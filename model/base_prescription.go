package model

import "time"

type BasePrescription struct {
	PrescriptionId     int       `json:"prescription_id"`
	OrgId              int       `json:"org_id"`
	UserId             string    `json:"user_id"`
	ConsultId          int       `json:"consult_id"`
	PrescriptionType   string    `json:"prescription_type"`
	DoctorId           string    `json:"doctor_id"`
	DoctorName         string    `json:"doctor_name"`
	CreateTime         time.Time `json:"create_time"`
	PrescriptionStatus string    `json:"prescription_status"`
}
