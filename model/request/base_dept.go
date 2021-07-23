package request

type AddDept struct {
	DeptName string `json:"dept_name"`
}

type FindDept struct {
	DeptName string `json:"dept_name"`
}