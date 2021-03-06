package handler

import (
	"net/http"
	"beautyfarm4market/util"
	"beautyfarm4market/entity"
	"encoding/json"
	"beautyfarm4market/dal"
	"strconv"
)

func PurchaseResHandler(w http.ResponseWriter, r *http.Request) {
	mappingOrderNo := r.FormValue("mappingOrderNo")
	tempOrderInfo:=dal.GetOrdersByMappingOrderNo(mappingOrderNo)
	indexUrl :=r.Host+"/?productId="+strconv.FormatInt(tempOrderInfo.ProductId,10)+"&channelcode="+tempOrderInfo.Channel
	if r.Method == "GET" {
		checkPurchaseResponse := getCheckPurchaseResponse(mappingOrderNo)
		checkPurchaseResponse.IndexUrl = indexUrl

		locals := make(map[string]interface{})
		locals["checkPurchaseResponse"] = checkPurchaseResponse
		util.RenderHtml(w, "purchaseRes.html", locals)
		return
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		mappingOrderNo := r.FormValue("mappingOrderNo")
		checkPurchaseResponse := getCheckPurchaseResponse(mappingOrderNo)
		checkPurchaseResponse.IndexUrl =indexUrl
		json.NewEncoder(w).Encode(checkPurchaseResponse)
		return
	}
	return
}

func getCheckPurchaseResponse(mappingOrderNo string) CheckPurchaseResponse {
	checkPurchaseResponse := CheckPurchaseResponse{}
	tempOrderInfo := dal.GetOrdersByMappingOrderNo(mappingOrderNo)
	if tempOrderInfo.MappingOrderNo != "" {
		checkPurchaseResponse.IsSucess = true
		checkPurchaseResponse.PayStatus = tempOrderInfo.PayStatus
		checkPurchaseResponse.MappingOrderNo=tempOrderInfo.MappingOrderNo
	} else {
		checkPurchaseResponse.IsSucess = true
		checkPurchaseResponse.Code = "400"
		checkPurchaseResponse.Message = "订单不存在"
	}
	return checkPurchaseResponse
}

type CheckPurchaseResponse struct {
	entity.BaseResultEntity
	IndexUrl       string `json:"indexUrl"`
	PayStatus      int  `json:"payStatus"` //支付状态 0未支付 1 支付中 2已支付  已退款
	MappingOrderNo string `json:"mappingOrderNo"`
}
