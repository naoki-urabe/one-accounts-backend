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

func getAllBanks(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var banks []models.Bank
	models.GetAllBanks(&banks)
	responseBody, err := json.Marshal(banks)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
}

func addBank(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	var bank models.Bank
	if err := json.Unmarshal(reqBody, &bank); err != nil {
		log.Fatal(err)
	}
	models.InsertBank(&bank)
	responseBody, err := json.Marshal(bank)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
}

func getDetails(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return
	}
	var details []models.Detail
	models.GetAccountDetails(&details)
	responseBody, err := json.Marshal(details)
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println(detail)
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
	router.HandleFunc("/api/banks", getAllBanks).Methods("GET")
	router.HandleFunc("/api/banks", addBank).Methods("POST","OPTIONS")
	fmt.Println("Listen 8080...")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
