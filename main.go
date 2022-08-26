package main

import (
	"dbssever/processes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	address := ":8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/dbsservers/authenticate", processes.AuthenticateUser).Methods("POST")
	router.HandleFunc("/dbsservers/associate", processes.AssociateUser).Methods("POST")
	router.HandleFunc("/dbsservers/otp/generate", processes.GenerateOtp).Methods("POST")
	router.HandleFunc("/dbsservers/otp/verify", processes.VerifyOtp).Methods("POST")
	log.Printf("main: starting the server on %v", address)
	log.Fatal(http.ListenAndServe(address, router))
}
