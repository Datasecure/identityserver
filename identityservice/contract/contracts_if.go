package contract

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ContractsInterface interface {
	// Create a new contract.
	// It is handler for POST /contracts
	Post(http.ResponseWriter, *http.Request)

	// Get a contract
	// It is handler for GET /contracts/{contractId}
	contractIdGet(http.ResponseWriter, *http.Request)

	// Sign a contract
	// It is handler for POST /contracts/{contractId}/signatures
	contractIdsignaturesPost(http.ResponseWriter, *http.Request)
}

func ContractsInterfaceRoutes(r *mux.Router, i ContractsInterface) {
	r.HandleFunc("/contracts", i.Post).Methods("POST")
	r.HandleFunc("/contracts/{contractId}", i.contractIdGet).Methods("GET")
	r.HandleFunc("/contracts/{contractId}/signatures", i.contractIdsignaturesPost).Methods("POST")
}
