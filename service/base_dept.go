package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"errors"
	"gorm.io/gorm"
)

func AddDept(dept model.BaseDept) error{
	if !errors.Is(global.MCS_DB.Where("dept_name = ?",dept.DeptName).First(&model.BaseDept{}).Error,gorm.ErrRecordNotFound){
		return errors.New("科室名称已存在，请修改科室名称")
	}
	return global.MCS_DB.Create(&dept).Error
}

func ListAllDept()(err error,depts []model.BaseDept)  {
	err = global.MCS_DB.Model(&model.BaseDept{}).Find(&depts).Error
	return
}

func DeleteDept(dept model.BaseDept) error {
	return global.MCS_DB.Where("dept_id = ?",dept.DeptId).Delete(&dept).Error
}

func UpdateDept(dept model.BaseDept) error {
	if !errors.Is(global.MCS_DB.Where("dept_name = ?",dept.DeptName).First(&model.BaseDept{}).Error,gorm.ErrRecordNotFound){
		return errors.New("科室名称已存在")
	}
	return global.MCS_DB.Where("dept_id = ?",dept.DeptId).Save(&dept).Error
}

func FindDept(name string) (err error,data []model.BaseDept){
	err = global.MCS_DB.Model(&model.BaseDept{}).Where("dept_name LIKE ?","%"+name+"%").Find(&data).Error
	return
}