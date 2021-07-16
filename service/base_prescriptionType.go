package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
)

func AddPrescription(prescription model.BasePrescription) error {
	return global.MCS_DB.Create(&prescription).Error
}