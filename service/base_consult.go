package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
)

func AddConsult(consult model.BaseConsult) error {
	return global.MCS_DB.Create(&consult).Error
}

func FindConsultByUser(userId string)(err error,consults []model.BaseConsult)  {
	err = global.MCS_DB.Model(&model.BaseConsult{}).Where("create_user_id = ?",userId).Find(&consults).Error
	return
}

func FindConsultByDoctor(doctorId string)(err error,consults []model.BaseConsult)  {
	err = global.MCS_DB.Model(&model.BaseConsult{}).Where("doctor_id = ?",doctorId).Find(&consults).Error
	return
}