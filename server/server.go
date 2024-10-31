package server

import (
	"MyStudy/order"
	"net/http"
)

func Server() {
	http.HandleFunc("/health", order.GetStatus)
	http.HandleFunc("/listOrder", order.GetDataBaseOrders)
	http.HandleFunc("/addOrder", order.AddOrder)
	http.HandleFunc("/confirmOrder", order.ConfirmOrder)
	http.HandleFunc("/cancelOrder", order.CancelOrder)

}
