package service

import (
	"MCS_Server/global"
	"MCS_Server/model"
	"MCS_Server/model/request"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


func Login(loginMsg request.Login) (err error, account model.BaseAccount) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + loginMsg.AppId + "&secret=" + loginMsg.AppSecret +
		"&js_code=" + loginMsg.Code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if strings.Contains(string(body), "openid") {
		requestBody := make(map[string]interface{})
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			return
		}
		opId := requestBody["openid"].(string)
		if !errors.Is(global.MCS_DB.Where("user_id = ? ", opId).First(&account).Error, gorm.ErrRecordNotFound) {
			return
		}else {
			account.UserId = opId
			account.UserType = "1"
			account.MiniOpenId = opId
			account.PhoneNo = loginMsg.Phone
			account.CreateTime = time.Now()
			err = global.MCS_DB.Create(&account).Error
			return
		}
	} else {
		err = errors.New("微信小程序没有返回对应openid")
		return
	}
}
