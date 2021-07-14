package request

type Login struct {
	Code string `json:"code"`
	PhoneNo string `json:"phone_no"`
}