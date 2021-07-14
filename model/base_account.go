package model

import "time"

type SysAccount struct {
	Code string
	PhoneNo string
	MiniOpenId string
	CreateTime time.Time
}