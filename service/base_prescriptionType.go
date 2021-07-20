package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"MCS_Server/model/response"
	"fmt"
)

func AddPrescription(prescription model.BasePrescription) error {
	return global.MCS_DB.Create(&prescription).Error
}

func AddPrescriptionDrug(drug model.BasePrescriptionDrug) error{
	return global.MCS_DB.Create(drug).Error
}

func ListPrescription(consultId int) (err error,data response.ListPrescription){
	var prescriptions []model.BasePrescription
	err = global.MCS_DB.Model(&model.BasePrescription{}).Where("consult_id = ?",consultId).Find(&prescriptions).Error
	if err != nil{
		return
	}
	for i:=0;i< len(prescriptions);i++{
		id := prescriptions[i].PrescriptionId
		var drugs []model.BasePrescriptionDrug
		prescription := response.Prescription{
			PrescriptionId: prescriptions[i].PrescriptionId,
			OrgId: prescriptions[i].OrgId,
			UserId: prescriptions[i].UserId,
			ConsultId: prescriptions[i].ConsultId,
			PrescriptionType: prescriptions[i].PrescriptionType,
			DoctorId: prescriptions[i].DoctorId,
			DoctorName: prescriptions[i].DoctorName,
			CreateTime: prescriptions[i].CreateTime,
			PrescriptionStatus: prescriptions[i].PrescriptionStatus,
			Drugs: make([]model.BasePrescriptionDrug,0),
		}
		err = global.MCS_DB.Model(&model.BasePrescriptionDrug{}).Where("prescription_drug_id = ?",id).Find(&drugs).Error
		if err != nil{
			return
		}
		for j:=0;j< len(drugs);j++ {
			modelDrug := drugs[j]
			prescription.Drugs = append(prescription.Drugs, modelDrug)
		}
		data.Prescriptions = append(data.Prescriptions, prescription)
	}
	return
}

func SubmitPrescription(ids []int) (err error){
	for i:=0;i<len(ids);i++ {
		fmt.Println(ids[i])
		err = global.MCS_DB.Model(&model.BasePrescription{}).Where("prescription_id = ?",ids[i]).Update("prescription_status","1").Error
		if err !=nil{
			return
		}
	}
	return
}
