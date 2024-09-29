package swagger

import (
	"fmt"
	"net/http"

	"github.com/fomichalopoulos/companiesMicroService/helpers"
	"github.com/fomichalopoulos/companiesMicroService/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("company name from GET is: ", companyName)

	company := &models.Company{}
	filter := bson.D{{Key: "name", Value: companyName}}
	client := MongoClient
	if err := helpers.FetchUsingProperty(
		client,
		"companyDB",
		"companies",
		filter,
		company,
		"company"); err != nil {
		if err == mongo.ErrNoDocuments {
			helpers.ErrorResponse(w, "No company with name = "+companyName+" exists in the database", http.StatusNotFound)
			return
		} else {
			helpers.ErrorResponse(w, "Fetching company database error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	helpers.RespondWithJSON(w, company)

}
