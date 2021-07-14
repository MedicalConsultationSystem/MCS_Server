package verify

var (
	LoginVerify  = Rules{"Code": {NotEmpty},"Phone":{NotEmpty}}
	AddOrgVerify = Rules{"OrgName": {NotEmpty}}
	AddDeptVerify = Rules{"DeptName":{NotEmpty}}
)
