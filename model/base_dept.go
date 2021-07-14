package model

type BaseDept struct {
	DeptId   int    `json:"dept_id"`
	DeptName string `json:"dept_name"`
}
func (m *BaseDept) TableName() string {
	return "base_dept"
}