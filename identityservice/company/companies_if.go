package company

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"github.com/gorilla/mux"
	"net/http"
)

type CompaniesInterface interface {
	// Register a new company
	// It is handler for POST /companies
	Post(http.ResponseWriter, *http.Request)

	// Update existing company. Updating ``globalId`` is not allowed.
	// It is handler for PUT /companies/{globalId}
	globalIdPut(http.ResponseWriter, *http.Request)

	// It is handler for GET /companies/{globalId}/validate
	globalIdvalidateGet(http.ResponseWriter, *http.Request)

	// Get the contracts where the organization is 1 of the parties. Order descending by
	// date.
	// It is handler for GET /companies/{globalId}/contracts
	globalIdcontractsGet(http.ResponseWriter, *http.Request)

	// It is handler for GET /companies/{globalId}/info
	globalIdinfoGet(http.ResponseWriter, *http.Request)
}

func CompaniesInterfaceRoutes(r *mux.Router, i CompaniesInterface) {
	r.HandleFunc("/companies", i.Post).Methods("POST")
	r.HandleFunc("/companies/{globalId}", i.globalIdPut).Methods("PUT")
	r.HandleFunc("/companies/{globalId}/validate", i.globalIdvalidateGet).Methods("GET")
	r.HandleFunc("/companies/{globalId}/contracts", i.globalIdcontractsGet).Methods("GET")
	r.HandleFunc("/companies/{globalId}/info", i.globalIdinfoGet).Methods("GET")
}
