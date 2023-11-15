package http

import (
	"log"
	"myapp/services"
	"net/http"
)

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/orders", ordersEndpoint)
}

func ordersEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	orders := query["order_uid"]
	log.Printf("Client is attempting to get orders by UID = %s", orders)

	if len(orders) > 0 {
		for _, uid := range orders {
			services.SelectOrderByUid(uid)
		}
	}
}

func StartListening(addr string) {
	log.Printf("Server is starting to listen at address '%s'...\n", addr)

	mux := http.NewServeMux()
	setupHandlers(mux)
	err := http.ListenAndServe(addr, mux)

	if err != nil {
		StopListening()
		log.Println(err)
		return
	}
}

func StopListening() {
	log.Printf("Server is stop listening...\n")
}
