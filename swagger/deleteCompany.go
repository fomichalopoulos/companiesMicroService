package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/fomichalopoulos/companiesMicroService/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func DelCompany(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	companyName := queryParams.Get("name")

	if companyName == "" {
		helpers.ErrorResponse(w, "Please provide a company name", http.StatusBadRequest)
		return
	}

	if len(companyName) < 1 || len(companyName) > 15 {
		helpers.ErrorResponse(w, "Please provide a valid company name of length between 1 to 15 characters", http.StatusBadRequest)
		return
	}

	fmt.Println("company name from DELETE is: ", companyName)

	if err := checkJWT(w, r); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	filter := bson.D{{Key: "name", Value: companyName}}
	client := MongoClient
	if err := helpers.DeleteOne(
		client,
		"companyDB",
		"companies",
		filter,
		"company"); err != nil {
		spew.Dump(err)
		if strings.Contains(err.Error(), "No company exists") {
			helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		} else {
			helpers.ErrorResponse(w, "Error in Deleting company: "+companyName+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	helpers.RespondWithJSON(w, "company with name = "+companyName+" deleted")

	// Produce an event to KAFKA
	eventMsg := "Company with name= " + companyName + " was deleted"
	if err := helpers.ProdKafkaMsg(eventMsg, kafkaConfig, kafkaProducer); err != nil {
		fmt.Println("error in kafka produce: ", err)
	}

}
