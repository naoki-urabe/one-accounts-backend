package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"one-accounts/models"
)

func getDetails(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return
	}
	vars := mux.Vars(r)
	responseBody, err := json.Marshal(vars["bank"])
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func addDetail(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var detail models.Detail
	if err := json.Unmarshal(reqBody, &detail); err != nil {
		log.Fatal(err)
	}
	models.InsertDetail(&detail)
	responseBody, err := json.Marshal(detail)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Content-Type", "application/json")
}

func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/accounts/{bank}/details", getDetails).Methods("GET")
	router.HandleFunc("/api/accounts/{bank}/details", addDetail).Methods("POST","OPTIONS")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
