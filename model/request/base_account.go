package request

type Login struct {
	//AppId string `json:"app_id"`
	//AppSecret string `json:"app_secret"`
	//Code string `json:"code"`
	OpenId string `json:"open_id"`
	Phone string `json:"phone"`
}