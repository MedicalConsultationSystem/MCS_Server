package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"errors"
	"gorm.io/gorm"
)

func AddOrg(org model.BaseOrg) error {
	if !errors.Is(global.MCS_DB.Where("org_name = ?",org.OrgName).First(&model.BaseOrg{}).Error,gorm.ErrRecordNotFound){
		return errors.New("机构名称已存在，请修改机构名称")
	}
	return global.MCS_DB.Create(&org).Error
}

func ListAllOrg() (err error,orgs []model.BaseOrg) {
	err = global.MCS_DB.Model(&model.BaseOrg{}).Find(&orgs).Error
	return
}

func DeleteOrg(org model.BaseOrg) error {
	return global.MCS_DB.Where("org_id = ?",org.OrgId).Delete(&org).Error
}

func UpdateOrg(org model.BaseOrg) error {
	if !errors.Is(global.MCS_DB.Where("org_name = ?",org.OrgName).First(&model.BaseOrg{}).Error,gorm.ErrRecordNotFound){
		return errors.New("机构名称已存在")
	}
	return global.MCS_DB.Where("org_id = ?",org.OrgId).Save(&org).Error
}