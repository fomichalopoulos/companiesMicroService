package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fomichalopoulos/companiesMicroService/models"
)

func DecodeJSON(r *http.Request, v interface{}) error {
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(v); err != nil {
		if errors.As(err, &unmarshalErr) {
			return fmt.Errorf("Invalid JSON format: %w", err)
		} else {
			return fmt.Errorf("Invalid JSON format: %w", err)
		}
	}
	return nil
}

var AllowedTypes = []string{"Corporations", "NonProfit", "Cooperative", "SoleProprietorship"}

func isValidType(typeValue string) bool {
	for _, allowedType := range AllowedTypes {
		if typeValue == allowedType {
			return true
		}
	}
	return false
}

func ValidateCompany4Patch(company models.Company) error {

	if company.Name != nil {
		companyName := *company.Name
		if len(companyName) < 1 || len(companyName) > 15 {
			return errors.New("Company name should not exceed 15 characters or be an empty string")
		}
	}
	description := company.Description
	if len(description) > 3000 {
		return errors.New("Description should not exceed 3000 characters")
	}
	if company.Type != nil {
		typeCompany := company.Type
		if !isValidType(*typeCompany) {
			return errors.New("Type Company is not one of required ones...")
		}
	}
	return nil

}

func ValidateCompany(company models.Company) error {

	/*
		if *company.Amount == 0 || *company.Type == "" {
			return errors.New("Please provide both name, amount, registered and type fields for a company")
		}
	*/
	if company.Name == nil || company.Amount == nil || company.Type == nil {
		return errors.New("Please provide both name, amount, registered and type fields for a company")
	}

	companyName := *company.Name
	if len(companyName) < 1 || len(companyName) > 15 {
		return errors.New("Company name should not exceed 15 characters or be an empty string")
	}
	description := company.Description
	if len(description) > 3000 {
		return errors.New("Description should not exceed 3000 characters")
	}

	typeCompany := company.Type
	if !isValidType(*typeCompany) {
		return errors.New("Type Company is not one of required ones...")
	}
	/*
		if company.Amount == nil || company.Registered == nil {
			return errors.New("Please provide both amount and registered fields for a company")
		}
	*/

	return nil

}
func RespondWithJSON(w http.ResponseWriter, data interface{}) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error in marshalling data: "+err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsonResp); err != nil {
		log.Printf("Error %v occured while sending response", err)
	}
}

func ErrorResponse(w http.ResponseWriter, message string, code int) {
	/*
		aMessage := messageError{
			Message: message,
		}
		jsonResp, err := json.Marshal(aMessage)
		if err != nil {
			http.Error(w, "Error in marshalling data: "+err.Error(), http.StatusInternalServerError)
		}
	*/
	//w.WriteHeader(http.StatusOK)
	w.WriteHeader(code)
	/*
		if _, err = w.Write(jsonResp); err != nil {
			log.Printf("Error %v occured while sending response", err)
		}
	*/
	fmt.Fprint(w, message)

}
