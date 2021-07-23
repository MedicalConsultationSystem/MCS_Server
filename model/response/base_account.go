package response

import "MCS_Server/model"

type AccountWithToken struct {
	model.BaseAccount
	Token string `json:"token"`
}
