package request

import "time"

type AddConsult struct {
	OrgId            int       `json:"org_id"`
	OrgName          string    `json:"org_name"`
	DeptId           int       `json:"dept_id"`
	DeptName         string    `json:"dept_name"`
	DoctorId         string    `json:"doctor_id"`
	DoctorName       string    `json:"doctor_name"`
	DoctorLevelCode  string    `json:"doctor_level_code"`
	DoctorLevelName  string    `json:"doctor_level_name"`
	CreateUserId     string    `json:"create_user_id"`
	PersonName       string    `json:"person_name"`
	PersonCardType   string    `json:"person_card_type"`
	PersonCardId     string    `json:"person_card_id"`
	PersonGenderCode string    `json:"person_gender_code"`
	PersonGenderName string    `json:"person_gender_name"`
	PersonBirthDate  time.Time `json:"person_birth_date"`
	PersonAge        int       `json:"person_age"`
	PersonPhoneNo    string    `json:"person_phone_no"`
	Question         string    `json:"question"`
	Diagnosis        string    `json:"diagnosis"`
	DrugIds          string    `json:"drug_ids"`
	DrugNames        string    `json:"drug_names"`
	PhotoIds         string    `json:"photo_ids"`
	ConsultStatus    int       `json:"consult_status"`
}

type FindConsultByUser struct {
	CreateUserId     string    `json:"create_user_id"`
}

type FindConsultByDoctor struct {
	DoctorId     string    `json:"doctor_id"`
}