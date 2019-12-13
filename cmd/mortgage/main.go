package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mortgage/cmd/mortgage/background"
	"mortgage/cmd/mortgage/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/request", controllers.CreateRequest).Methods("POST")
	router.HandleFunc("/request", controllers.GetRequests).Methods("GET")
	router.HandleFunc("/request/{id}", controllers.GetStatusByRequestId).Methods("GET")

	port := os.Getenv("PORT")

	background.StartBackgroundTasks()

	log.Println(fmt.Sprintf("Backend strated. Port listen %s", port))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
