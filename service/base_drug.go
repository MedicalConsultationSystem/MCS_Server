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