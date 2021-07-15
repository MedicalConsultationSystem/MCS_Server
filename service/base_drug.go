package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
)

func ListAllDrug() (err error, drugs []model.BaseDrug) {
	err = global.MCS_DB.Model(&model.BaseDrug{}).Find(&drugs).Error
	return
}
