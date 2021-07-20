package verify

var (
	LoginVerify   = Rules{"Code": {NotEmpty}, "Phone": {NotEmpty}}
	AddOrgVerify  = Rules{"OrgName": {NotEmpty}}
	AddDeptVerify = Rules{"DeptName": {NotEmpty}}
	AddDrugVerify = Rules{"DrugName": {NotEmpty}, "TradeName": {NotEmpty},
		"PinyinCode": {NotEmpty}, "Specification": {NotEmpty}, "PackUnit": {NotEmpty},
		"Price": {NotEmpty}, "Dose": {NotEmpty}, "DoseUnit": {NotEmpty}}
	AddDoctorVerify = Rules{"DoctorId":{NotEmpty},"DoctorName":{NotEmpty},
		"OrgId":{NotEmpty},"OrgName":{NotEmpty},"DeptId":{NotEmpty},"DeptName":{NotEmpty}}
	BaseDeptVerify = Rules{"DeptId":{NotEmpty},"DeptName":{NotEmpty}}
	BaseOrgVerify = Rules{"OrgId":{NotEmpty},"OrgName":{NotEmpty}}
	BaseDrugVerify = Rules{"DrugId":{NotEmpty},"DrugName": {NotEmpty}, "TradeName": {NotEmpty},
		"PinyinCode": {NotEmpty}, "Specification": {NotEmpty}, "PackUnit": {NotEmpty},
		"Price": {NotEmpty}, "Dose": {NotEmpty}, "DoseUnit": {NotEmpty}}
	AddConsultVerify = Rules{"OrgId":{NotEmpty},"OrgName":{NotEmpty},"DeptId":{NotEmpty},
		"DeptName":{NotEmpty},"DoctorId":{NotEmpty},"DoctorName":{NotEmpty},
		"CreateUserId":{NotEmpty},"PersonName":{NotEmpty},"PersonCardType":{NotEmpty},
		"PersonCardId":{NotEmpty},"PersonGenderCode":{NotEmpty},"PersonGenderName":{NotEmpty},
		"PersonBirthDate":{NotEmpty},"PersonAge":{NotEmpty},"PersonPhoneNo":{NotEmpty},
		"Question":{NotEmpty},"Diagnosis":{NotEmpty},"DrugIds":{NotEmpty},"DrugNames":{NotEmpty}}
	AddPrescriptionVerify = Rules{"OrgId": {NotEmpty},"UserId":{NotEmpty},"ConsultId":{NotEmpty},
		"PrescriptionType":{NotEmpty},"DoctorId":{NotEmpty},"DoctorName":{NotEmpty}}
	ListPrescriptionVerify = Rules{"ConsultId":{NotEmpty}}
	AddPrescriptionDrugVerify = Rules{"OrgId":{NotEmpty},"PrescriptionId":{NotEmpty}}
)
