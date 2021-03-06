package request

type AddDoctor struct {
	DoctorId   string `json:"doctor_id"`
	DoctorName string `json:"doctor_name"`
	OrgId      int    `json:"org_id"`
	OrgName    string `json:"org_name"`
	DeptId     int    `json:"dept_id"`
	DeptName   string `json:"dept_name"`
	AvatarUrl  string `json:"avatar_url"`
	LevelCode  string `json:"level_code"`
	LevelName  string `json:"level_name"`
}

type FindDoctorByName struct {
	DoctorName string `json:"doctor_name"`
}

type DeleteDoctor struct {
	DoctorId   string `json:"doctor_id"`
}