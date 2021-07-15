package verify

var (
	LoginVerify   = Rules{"Code": {NotEmpty}, "Phone": {NotEmpty}}
	AddOrgVerify  = Rules{"OrgName": {NotEmpty}}
	AddDeptVerify = Rules{"DeptName": {NotEmpty}}
	AddDrugVerify = Rules{"DrugName": {NotEmpty}, "TradeName": {NotEmpty},
		"PinyinCode": {NotEmpty}, "Specification": {NotEmpty}, "PackUnit": {NotEmpty},
		"Price": {NotEmpty}, "Dose": {NotEmpty}, "DoseUnit": {NotEmpty}}
)
