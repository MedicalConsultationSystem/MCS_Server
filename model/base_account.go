package model

import "time"

type BaseAccount struct {
	Code string
	PhoneNo string
	MiniOpenId string
	CreateTime time.Time
}