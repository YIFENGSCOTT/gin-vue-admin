package EncryptionKey

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/EncryptionKey"
	EncryptionKeyReq "github.com/flipped-aurora/gin-vue-admin/server/model/EncryptionKey/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type EncryptionKeyApi struct {
}

var encryptionKeyService = service.ServiceGroupApp.EncryptionKeyServiceGroup.EncryptionKeyService

type rollData struct {
	QrCodeUrl string
	Content string
	Type int
	QrCodeBase64 bool
}

type postResp struct {
	Code int
	Msg string
	Data rollData
}

// CreateEncryptionKey 创建EncryptionKey
// @Tags EncryptionKey
// @Summary 创建EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EncryptionKey.EncryptionKey true "创建EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /encryptionKey/createEncryptionKey [post]
func (encryptionKeyApi *EncryptionKeyApi) CreateEncryptionKey(c *gin.Context) {
	var encryptionKey EncryptionKey.EncryptionKey
	_ = c.ShouldBindJSON(&encryptionKey)
	keyContent := "EK-" + strconv.FormatInt(int64(rand.Int()), 10)

	//您的秘钥信息如下：
	app_id := "kfofoph1fl3vfcen"
	app_secret := "aWVJVGlPM2x1dWlZbjNxTFlRZDZ3dz09"

	requestString := "https://www.mxnzp.com/api/qrcode/create/single?content=" +
		keyContent +
		"&size=500&type=0&" +
		"app_id=" +
		app_id +
		"&app_secret=" +
		app_secret

	encryptionKey.KeyIllustrationUrl = requestString

	client := http.Client{}
	req,_ := http.NewRequest("GET", requestString,nil)
	resp,_ := client.Do(req)

	body,_ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	respMap := postResp{}
	err := json.Unmarshal([]byte(bodyStr),&respMap)

	if err != nil{
		return
	}

	encryptionKey.KeyContent = keyContent
	if err := encryptionKeyService.CreateEncryptionKey(encryptionKey); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}



	mailMsg := "Hello there, <strong>Distrust</strong> recently received a request to set up your new encryption key for distrust.yifengsun.com: \n\n" +
		"<h1 style=3D\"font-weight:700;word-break:break-word;font-size:46px;line-height:52px\"><code>" + keyContent + "</code>\n</h1>" +
		"<img src=\""+ respMap.Data.QrCodeUrl +"\" alt=\"illustration\" height=\"120\" width=\"120\">"+
		"<div>Please use this key to start a thread ;) \n\n</div>"+
		"<div><strong>Thank you for participating in my graduation design/defense,</strong></div>"+
		"<div>Sincerely,</div>"+
		"<div>Yifeng Sun from Distrust</div>"
	m := gomail.NewMessage()
	m.SetHeader("From", "Distrust@sunyifeng.buzz")
	m.SetHeader("To", encryptionKey.Beiyong)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Your new encryption key from distrust.yifengsun.com")
	m.SetBody("text/html", mailMsg)
	//m.Attach("/home/Alex/lolcat.jpg")



	d := gomail.NewDialer("smtp.ym.163.com", 25, "distrust@sunyifeng.buzz", "Di42419629")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// DeleteEncryptionKey 删除EncryptionKey
// @Tags EncryptionKey
// @Summary 删除EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EncryptionKey.EncryptionKey true "删除EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /encryptionKey/deleteEncryptionKey [delete]
func (encryptionKeyApi *EncryptionKeyApi) DeleteEncryptionKey(c *gin.Context) {
	var encryptionKey EncryptionKey.EncryptionKey
	_ = c.ShouldBindJSON(&encryptionKey)
	if err := encryptionKeyService.DeleteEncryptionKey(encryptionKey); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEncryptionKeyByIds 批量删除EncryptionKey
// @Tags EncryptionKey
// @Summary 批量删除EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /encryptionKey/deleteEncryptionKeyByIds [delete]
func (encryptionKeyApi *EncryptionKeyApi) DeleteEncryptionKeyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := encryptionKeyService.DeleteEncryptionKeyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEncryptionKey 更新EncryptionKey
// @Tags EncryptionKey
// @Summary 更新EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body EncryptionKey.EncryptionKey true "更新EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /encryptionKey/updateEncryptionKey [put]
func (encryptionKeyApi *EncryptionKeyApi) UpdateEncryptionKey(c *gin.Context) {
	var encryptionKey EncryptionKey.EncryptionKey
	_ = c.ShouldBindJSON(&encryptionKey)
	if err := encryptionKeyService.UpdateEncryptionKey(encryptionKey); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEncryptionKey 用id查询EncryptionKey
// @Tags EncryptionKey
// @Summary 用id查询EncryptionKey
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EncryptionKey.EncryptionKey true "用id查询EncryptionKey"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /encryptionKey/findEncryptionKey [get]
func (encryptionKeyApi *EncryptionKeyApi) FindEncryptionKey(c *gin.Context) {
	var encryptionKey EncryptionKey.EncryptionKey
	_ = c.ShouldBindQuery(&encryptionKey)
	if err, reencryptionKey := encryptionKeyService.GetEncryptionKey(encryptionKey.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reencryptionKey": reencryptionKey}, c)
	}
}

// GetEncryptionKeyList 分页获取EncryptionKey列表
// @Tags EncryptionKey
// @Summary 分页获取EncryptionKey列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EncryptionKeyReq.EncryptionKeySearch true "分页获取EncryptionKey列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /encryptionKey/getEncryptionKeyList [get]
func (encryptionKeyApi *EncryptionKeyApi) GetEncryptionKeyList(c *gin.Context) {
	var pageInfo EncryptionKeyReq.EncryptionKeySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := encryptionKeyService.GetEncryptionKeyInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
