package model

import "time"

type BaseAccount struct {
	UserId     string    `json:"user_id"`
	UserType   string    `json:"user_type"`
	MiniOpenId string    `json:"mini_open_id"`
	PhoneNo    string    `json:"phone_no"`
	CreateTime time.Time `json:"create_time"`
}
func (m *BaseAccount) TableName() string {
	return "base_account"
}
