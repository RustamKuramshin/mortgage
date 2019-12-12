package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/request", CreateRequest).Methods("POST")
	router.HandleFunc("/request/{id}", GetStatusByRequestId).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}

type RequestResponse struct {
	Id         string `json:"id"`
	StatusCode string `json:"status_code"`
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	uuid, _ := uuid.NewV4()
	json.NewEncoder(w).Encode(RequestResponse{Id: uuid.String(), StatusCode: "processing"})
}

func GetStatusByRequestId(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Check Status")
	id := mux.Vars(r)["id"]
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(RequestResponse{Id: id, StatusCode: getRandomStatus()})
}

func getRandomStatus() string {
	statuses := []string{"processing", "approved", "rejected", "error"}
	ri := rand.Intn(len(statuses))
	return statuses[ri]
}
