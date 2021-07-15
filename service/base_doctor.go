package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"gorm.io/gorm"
	"errors"
)

func ListAllDoctor() (err error, doctors []model.BaseDoctor) {
	err = global.MCS_DB.Model(&model.BaseDoctor{}).Find(&doctors).Error
	return
}

func AddDoctor(doctor model.BaseDoctor) error {
	if !errors.Is(global.MCS_DB.Where("doctor_id = ?",doctor.DoctorId).First(&model.BaseDoctor{}).Error,gorm.ErrRecordNotFound){
		return errors.New("医生已注册")
	}
	return global.MCS_DB.Create(&doctor).Error
}

func FindDoctorByName(doctorName string) (err error,doctors []model.BaseDoctor){
	err = global.MCS_DB.Model(&model.BaseDoctor{}).Where("doctor_name LIKE ?","%"+doctorName+"%").Find(&doctors).Error
	return
}