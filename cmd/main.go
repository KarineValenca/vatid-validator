package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/KarineValenca/vatid-validator/pkg/validator"
	"github.com/KarineValenca/vatid-validator/pkg/validator/soap"
	"github.com/gorilla/mux"
)

const (
	checkVatServiceURL = "https://ec.europa.eu/taxation_customs/vies/services/checkVatService"
)

type handler struct {
	validator validator.ValidatorApp
}

func main() {
	// check arguments to start program
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Failed to start program, invalid arguments")
	}

	// read config
	cfg := HTTPConfig{}
	if err := cfg.ReadConfig(args[1]); err != nil {
		log.Fatal("Failed to read config")
	}

	soapClient := soap.Client{
		Client:             http.DefaultClient,
		CheckVatServiceURL: checkVatServiceURL,
	}

	// handlers
	h := handler{
		validator: &soapClient,
	}

	r := mux.NewRouter()
	r.HandleFunc("/valid_vat_id", h.ValidateVatID)

	// http server
	server := &http.Server{
		Handler:      r,
		Addr:         strings.Join([]string{cfg.Addr, cfg.Port}, ":"),
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to serve http server")
	}
}
