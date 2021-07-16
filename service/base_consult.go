package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
)

func AddConsult(consult model.BaseConsult) error {
	return global.MCS_DB.Create(&consult).Error
}