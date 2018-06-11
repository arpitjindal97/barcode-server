package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func StartProcess() {

	muxHttp := mux.NewRouter()

	muxHttp.HandleFunc("/Status", StatusHandler).Methods("GET")
	muxHttp.HandleFunc("/Result", ResultHandler).Methods("POST")

	srv := &http.Server{
		Handler:      muxHttp,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Running")
}
func ResultHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	strBody := string(body)

	fmt.Fprint(w, strBody)
}
