package request

type Login struct {
	OpenId string `json:"open_id"`
	Phone string `json:"phone"`
}