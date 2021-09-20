package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KarineValenca/vatid-validator/pkg/validator/dto"
)

func (h *handler) ValidateVatID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Error: Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request dto.ValidateVATIDRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error: invalid request %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, err := h.validator.CheckValidVAT(request.CountryCode, request.VatNumber)
	if err != nil {
		log.Printf("Error: something went wrong checking VAT %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(dto.ValidateVATIDResponse{
		IsValid: isValid,
	})
	if err != nil {
		log.Printf("Error: something went wrong while marshal response %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
