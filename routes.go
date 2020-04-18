package main

import (
	"github.com/gorilla/mux"
	"github.com/p-weisk/secret-form/rest"
)

func registerRoutes(r *mux.Router) {

// Form result phase
	// view form result
	r.HandleFunc("/api/form/{fid}/result", rest.FormResult).Methods("GET")

// Form answering phase
	// add answer
	r.HandleFunc("/api/form/{fid}/answers", rest.AddAnswer).Methods("POST")
	// get empty form
	r.HandleFunc("/api/form/{fid}", rest.GetForm).Methods("GET")

// Form creation phase
	// create new form
	r.HandleFunc("/api/form", rest.NewForm).Methods("POST")

}
