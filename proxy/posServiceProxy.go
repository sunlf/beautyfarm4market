package proxy

import (
	"beautyfarm4market/entity"
	"fmt"
	"beautyfarm4market/config"
	"time"
	"beautyfarm4market/util"
)

type SoaIsVipRes struct {
	Status  int     `json:"status"`
	Message int     `json:"message"`
	Data    VipData `json:"data"`
}

type VipData struct {
	IsVip       bool `json:"isVip"`
	IsMarketVip bool `json:"IsMarketVip"`
}

//    "customer/isVipMember?org_no=beautyfarm&mobile=%s&appid=%s&timestamp=%s&sign=%s",
//serverRes http请求结果
func IsVip(mobileNo string) (soaIsVipResOut SoaIsVipRes, serverRes entity.BaseResultEntity) {
	var soaIsVipRes SoaIsVipRes
	methodUrl := fmt.Sprintf(config.ConfigInfo.PosService, config.ConfigInfo.IsVipUrl)

	appId := config.ConfigInfo.PosServiceAppId
	timeSpan := time.Now().Unix()
	sign := getSign(appId, timeSpan)

	url := fmt.Sprintf(methodUrl, mobileNo, config.ConfigInfo.PosServiceAppId, timeSpan, sign)
	serverRes = httpGetProxy(url, &soaIsVipRes)
	soaIsVipResOut = soaIsVipRes
	return
}

//   "appid=%s&secretkey=%s&timestamp=%s",
func getSign(appId string, timeSpan int64) string {
	original := fmt.Sprintf(config.ConfigInfo.SignTemplate, appId, config.ConfigInfo.AppSecret, timeSpan)
	return util.GetMd5(original)
}


