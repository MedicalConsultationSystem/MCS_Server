package request

type Login struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Code string `json:"code"`
	Phone string `json:"phone"`
}