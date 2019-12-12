package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mortgage/cmd/mortgage/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/request", controllers.CreateRequest).Methods("POST")
	router.HandleFunc("/request/{id}", controllers.GetStatusByRequestId).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
