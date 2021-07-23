package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"time"
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

func AcceptConsult(consultId int)(err error){
	err = global.MCS_DB.Model(&model.BaseConsult{}).Where("consult_id = ?",consultId).Update("consult_status",2).Update("accept_time",time.Now()).Error
	return
}

func FinishConsult(consultId int)(err error){
	err = global.MCS_DB.Model(&model.BaseConsult{}).Where("consult_id = ?",consultId).Update("consult_status",3).Update("finish_time",time.Now()).Error
	return
}