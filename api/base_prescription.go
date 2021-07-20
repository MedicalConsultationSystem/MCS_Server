package api

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"MCS_Server/model/request"
	"MCS_Server/model/response"
	"MCS_Server/service"
	"MCS_Server/utils/verify"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// @Tags 处方
// @Summary 增加一个处方
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddPrescription true "机构id，用户id，问诊信息id，处方类型，医生id，医生名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方新增成功"}"
// @Router /prescription/add [post]
func AddPrescription(c *gin.Context) {
	var newPrescriptionMsg request.AddPrescription
	_ = c.ShouldBindJSON(&newPrescriptionMsg)
	if err := verify.Verify(newPrescriptionMsg, verify.AddPrescriptionVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	prescription := model.BasePrescription{OrgId: newPrescriptionMsg.OrgId, UserId: newPrescriptionMsg.UserId,
		ConsultId: newPrescriptionMsg.ConsultId, PrescriptionType: newPrescriptionMsg.PrescriptionType, DoctorId: newPrescriptionMsg.DoctorId,
		DoctorName: newPrescriptionMsg.DoctorName, CreateTime: time.Now(), PrescriptionStatus: "0"}
	if err := service.AddPrescription(prescription); err != nil {
		global.MCS_Log.Error("处方新增失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		response.SuccessWithMsg("处方新增成功", c)
	}
}

// @Tags 处方
// @Summary 提交处方
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SubmitPrescription true "处方id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方提交成功"}"
// @Router /prescription/submit [post]
func SubmitPrescription(c *gin.Context) {
	var msg = request.SubmitPrescription{}
	_ = c.ShouldBindJSON(&msg)
	if err := verify.Verify(msg, verify.SubmitPrescriptionVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err := service.SubmitPrescription(msg.PrescriptionIds); err != nil {
		global.MCS_Log.Error("提交处方失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	}
	response.SuccessWithMsg("处方提交成功", c)

	//response.SuccessWithMsg("接口还没写完，让你失望了", c)
}

// @Tags 处方
// @Summary 查询处方信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ListPrescription true "问诊信息 id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方查询成功"}"
// @Router /prescription/list [post]
func ListPrescription(c *gin.Context) {
	var msg request.ListPrescription
	_ = c.ShouldBindJSON(&msg)
	if err := verify.Verify(msg, verify.ListPrescriptionVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if err, data := service.ListPrescription(msg.ConsultId); err != nil {
		global.MCS_Log.Error("处方信息获取失败", zap.Any("err", err))
		response.FailWithMsg("处方信息获取失败:"+err.Error(), c)
	} else {
		response.SuccessWithAll(data, "处方信息获取成功", c)
	}
}

// @Tags 处方
// @Summary 为某个处方添加一种药物
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PrescriptionDrug true "包含处方id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方药物添加成功"}"
// @Router /prescription/addDrug [post]
func AddPrescriptionDrug(c *gin.Context) {
	var newPreDrugMsg request.PrescriptionDrug
	_ = c.ShouldBindJSON(&newPreDrugMsg)
	fmt.Println(newPreDrugMsg)
	if err := verify.Verify(newPreDrugMsg, verify.AddPrescriptionDrugVerify); err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	drug := model.BasePrescriptionDrug{
		OrgId:          newPreDrugMsg.OrgId,
		PrescriptionId: newPreDrugMsg.PrescriptionId,
		DrugId:         newPreDrugMsg.DrugId,
		DrugName:       newPreDrugMsg.DrugName,
		Specification:  newPreDrugMsg.Specification,
		Dose:           newPreDrugMsg.Dose,
		DoseUnit:       newPreDrugMsg.DoseUnit,
		FrequencyCode:  newPreDrugMsg.FrequencyCode,
		FrequencyName:  newPreDrugMsg.FrequencyName,
		UsageCode:      newPreDrugMsg.UsageCode,
		UsageName:      newPreDrugMsg.UsageName,
		TakeDays:       newPreDrugMsg.TakeDays,
		Quantity:       newPreDrugMsg.Quantity,
		Price:          newPreDrugMsg.Price,
		PackUnit:       newPreDrugMsg.PackUnit,
		GroupNumber:    newPreDrugMsg.GroupNumber,
		SortNumber:     newPreDrugMsg.SortNumber,
		Remark:         newPreDrugMsg.Remark,
	}
	if err := service.AddPrescriptionDrug(drug); err != nil {
		global.MCS_Log.Error("处方药物添加失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		response.SuccessWithMsg("处方药物添加成功", c)
	}

}

// @Tags 处方
// @Summary 为某个处方删除一种药物
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeletePrescription true "包含处方药物id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"处方药物删除成功"}"
// @Router /prescription/delDrug [delete]
func DeletePrescriptionDrug(c *gin.Context) {
	var newPreDrugMsg request.DeletePrescription
	_ = c.ShouldBindJSON(&newPreDrugMsg)
	if err := service.DeletePrescriptionDrug(newPreDrugMsg.PrescriptionDrugId); err != nil {
		global.MCS_Log.Error("处方药物删除失败", zap.Any("err", err))
		response.FailWithMsg(err.Error(), c)
	} else {
		response.SuccessWithMsg("处方药物删除成功", c)
	}

}
