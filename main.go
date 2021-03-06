package main

import (
	"beautyfarm4market/handler"
	"net/http"
	"log"
	_ "beautyfarm4market/config"
	"beautyfarm4market/config"
)

func main() {
	cfg := config.GetConfigFromXml()
	mux := http.NewServeMux()
	handler.StaticDirHandler(mux, "/assets/", 0)
	mux.HandleFunc("/", handler.SafeHandler(handler.IndexHandler))
	mux.HandleFunc("/upload", handler.SafeHandler(handler.UploadHandler))
	mux.HandleFunc("/view", handler.SafeHandler(handler.ViewHandler))
	mux.HandleFunc("/list", handler.SafeHandler(handler.ListHandler))
	mux.HandleFunc("/sendMsg", handler.SafeHandler(handler.MessageHandler))
	mux.HandleFunc("/addOrder", handler.SafeHandler(handler.AddOrderHandler))
	mux.HandleFunc("/orderList", handler.SafeHandler(handler.OrderListHandler))
	mux.HandleFunc("/promotion", handler.SafeHandler(handler.RouteHandler))
	mux.HandleFunc("/payCallBack", handler.SafeHandler(handler.PayCallBackHandler))
	mux.HandleFunc("/purchaseRes", handler.SafeHandler(handler.PurchaseResHandler))
	mux.HandleFunc("/checkPurchaseRes", handler.SafeHandler(handler.PurchaseResHandler))
	mux.HandleFunc("/handlerWeChatLogin", handler.SafeHandler(handler.HandlerWeChatLoginHandler))
	mux.HandleFunc("/favicon.ico", handler.SafeHandler(handler.DefaultHandler))
	mux.HandleFunc("/prodcut", handler.SafeHandler(handler.ProductHandler))
	mux.HandleFunc("/prodcutdetail", handler.SafeHandler(handler.ProductDetailHandler))
	mux.HandleFunc("/report", handler.SafeHandler(handler.ReportHandler))
	mux.HandleFunc("/backyard", handler.SafeHandler(handler.BackyardHandler))
	mux.HandleFunc("/cancelOrder", handler.SafeHandler(handler.CancelOrderHandler))
	mux.HandleFunc("/refundOrder", handler.SafeHandler(handler.RefundOrderHandler))
	err := http.ListenAndServe(cfg.Port, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
