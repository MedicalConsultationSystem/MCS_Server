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
	var account model.BaseAccount
	if !errors.Is(global.MCS_DB.Where("phone_no = ?",doctor.DoctorId).First(&account).Error,gorm.ErrRecordNotFound){
		err :=global.MCS_DB.Model(&model.BaseAccount{}).Where("phone_no = ?",doctor.DoctorId).Update("user_type","2").Error
		if err != nil{
			return err
		}
	}
	return global.MCS_DB.Create(&doctor).Error
}

func FindDoctorByName(doctorName string) (err error,doctors []model.BaseDoctor){
	err = global.MCS_DB.Model(&model.BaseDoctor{}).Where("doctor_name LIKE ?","%"+doctorName+"%").Find(&doctors).Error
	return
}

func UpdateDoctor(doctor model.BaseDoctor) error {
	return global.MCS_DB.Where("doctor_id = ?",doctor.DoctorId).Save(&doctor).Error
}

func DeleteDoctor(doctorId string) (err error){
	var account model.BaseAccount
	err = global.MCS_DB.Where("phone_no = ?",doctorId).First(&account).Error
	if err !=nil{
		return
	}
	err =global.MCS_DB.Model(&model.BaseAccount{}).Where("phone_no = ?",doctorId).Update("user_type","1").Error
	if err != nil{
		return
	}
	err = global.MCS_DB.Where("doctor_id = ?",doctorId).Delete(&model.BaseDoctor{}).Error
	return
}