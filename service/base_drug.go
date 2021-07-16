package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"errors"
	"gorm.io/gorm"
)

func ListAllDrug() (err error, drugs []model.BaseDrug) {
	err = global.MCS_DB.Model(&model.BaseDrug{}).Find(&drugs).Error
	return
}

func AddDrug(drug model.BaseDrug) error {
	if !errors.Is(global.MCS_DB.Where("drug_name = ? or trade_name = ? or pinyin_code = ?",drug.DrugName,drug.TradeName,drug.PinyinCode).First(&model.BaseDrug{}).Error,gorm.ErrRecordNotFound){
		return errors.New("药物冲突，已存在")
	}
	return global.MCS_DB.Create(&drug).Error
}

func FindDrugByPinyin(pinyin string) (err error, drugs []model.BaseDrug) {
	err = global.MCS_DB.Model(&model.BaseDrug{}).Where("pinyin_code LIKE ?","%"+pinyin+"%").Find(&drugs).Error
	return
}

func FindDrugByName(drugName string)(err error, drugs []model.BaseDrug) {
	err = global.MCS_DB.Model(&model.BaseDrug{}).Where("drug_name LIKE ?","%"+drugName+"%").Find(&drugs).Error
	return
}

func DeleteDrug(drug model.BaseDrug) error {
	return global.MCS_DB.Where("drug_id = ?",drug.DrugId).Delete(&drug).Error
}

func ListAllUsage() (err error,usages []model.BaseUsage) {
	err = global.MCS_DB.Model(&model.BaseUsage{}).Find(&usages).Error
	return
}

func ListAllFrequency() (err error,frequencies []model.BaseFrequency){
	err = global.MCS_DB.Model(&model.BaseFrequency{}).Find(&frequencies).Error
	return
}