package http

import (
	"encoding/json"
	"log"
	db "myapp/database"
	"myapp/services"
	"net/http"
)

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/orders", ordersHandler)
}

func ordersHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, access-control-allow-methods, access-control-allow-credentials")
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")

	query := request.URL.Query()
	order_uids := query["order_uid"]

	if len(order_uids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Client is attempting to get orders by UID = %s", order_uids)

	foundContent := []*db.Order{}
	for _, uid := range order_uids {
		if order, err := services.SelectOrderByUID(uid); err == nil {
			foundContent = append(foundContent, order)
		}
	}

	if len(foundContent) > 0 {
		_, err := json.Marshal(foundContent)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		json.NewEncoder(w).Encode(foundContent)
	} else {
		w.WriteHeader(http.StatusNoContent)
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
