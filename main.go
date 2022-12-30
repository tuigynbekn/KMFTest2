package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handler := NewHandler()

	router.HandleFunc("/proxy", handler.proxyHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/history", handler.historyHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/history/{id}", handler.historyHandler).Methods(http.MethodGet, http.MethodOptions)

	http.Handle("/", router)

	fmt.Println("Server is listening... http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
