package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fomichalopoulos/companiesMicroService/helpers"
	"github.com/fomichalopoulos/companiesMicroService/models"
	"go.mongodb.org/mongo-driver/bson"
)

func PatchCompany(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	companyName := queryParams.Get("name")

	if err := checkJWT(w, r); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if companyName == "" {
		helpers.ErrorResponse(w, "Please provide a company name", http.StatusBadRequest)
		return
	}

	if len(companyName) < 1 || len(companyName) > 15 {
		helpers.ErrorResponse(w, "Please provide a valid company name of length between 1 to 15 characters", http.StatusBadRequest)
		return
	}

	fmt.Println("company name from PUT is: ", companyName)

	var company models.Company

	if err := helpers.DecodeJSON(r, &company); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helpers.ValidateCompany4Patch(company); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	bsonToPatch := bson.M{}
	if company.Name != nil {
		bsonToPatch["name"] = *company.Name
	}
	if company.Description != "" {
		bsonToPatch["description"] = *&company.Description
	}
	if company.Amount != nil {
		bsonToPatch["amount"] = *company.Amount
	}
	if company.Registered != nil {
		bsonToPatch["registered"] = *company.Registered
	}
	if company.Type != nil {
		bsonToPatch["type"] = *company.Type
	}

	filter := bson.M{"name": companyName}
	updateSet := bson.M{
		"$set": bsonToPatch,
		/*
			"$set": bson.M{
				"name":        company.Name,
				"description": company.Description,
				"amount":      company.Amount,
				"registered":  company.Registered,
				"type":        company.Type,
			},
		*/
	}
	client := MongoClient
	if err := helpers.UpdateUsingFilter(
		client,
		filter,
		"companyDB",
		updateSet,
		"companies"); err != nil {
		if strings.Contains(err.Error(), "No company was updated") {
			helpers.ErrorResponse(w, "No company was updated", http.StatusNotModified)
			return
		} else {
			helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	helpers.RespondWithJSON(w, "company "+*company.Name+" has been updated")

	// Produce an event to KAFKA
	eventMsg := "Company with name= " + *company.Name + " was patched"
	if err := helpers.ProdKafkaMsg(eventMsg, kafkaConfig, kafkaProducer); err != nil {
		fmt.Println("error in kafka produce: ", err)
	}

}
