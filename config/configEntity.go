package config

import "fmt"

var ConfigInfo configEntity

type configEntity struct {
	SmsUrl string
	OrderServiceUrl string
	AddOrderUrl string
	AccountRegisterUrl string
	OrderServiceAppId string
	AccountNo string
	Channel string
	SmsOfOrderSucess string
	GetOrderDetailUrl string
	SmsOfVaild string
	MobileCookie string
	TimeLayout string
	RegisterChannelType string
}

func init()  {
	ConfigInfo= configEntity{
		SmsUrl:"http://esms10.10690007.net/sms/mt",
		//下单接口配置
		OrderServiceUrl:"http://180.169.107.155:8866/api/v1/bf-cam-adapter/spec/%s",
		AddOrderUrl:"order/orderInsertOrUpdate",
		GetOrderDetailUrl:"order/orderDetailSelect?orderList=%s&appId=%s",
		AccountRegisterUrl:"account/accountRegister",
		OrderServiceAppId:"TEST",
		AccountNo:"129147",
		Channel:"XS0001",
		//下单接口配置END
		SmsOfOrderSucess:"您已成功购买产品%s，您的院余号为%s",
		SmsOfVaild:"%s（美丽田园手机验证码，请完成验证）， 如非本人操作，请忽略本短信。",
		MobileCookie:"code%s",
		TimeLayout:"2006-01-02 15:04:05",
		RegisterChannelType:"3003",
	}
	fmt.Printf("init Config")
}


