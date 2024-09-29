package swagger

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/fomichalopoulos/companiesMicroService/helpers"
	"github.com/fomichalopoulos/companiesMicroService/models"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtSecret = []byte("your_secret_key")

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mongo client = ", MongoClient)
	client := MongoClient

	if err := checkJWT(w, r); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var company models.Company
	if err := helpers.DecodeJSON(r, &company); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("company!!!!")
	spew.Dump(company)

	if err := helpers.ValidateCompany(company); err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if a company with the same name already exists in the MongoDB
	filter := bson.D{{Key: "name", Value: company.Name}}
	aCompany := &models.Company{}
	if err := helpers.CheckIfExists(
		client,
		"companyDB",
		"companies",
		filter,
		aCompany,
		"company"); err != nil {
		fmt.Printf("err = %v", err.Error())
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
		//return nil
	}
	fmt.Println("After Check if exists")
	spew.Dump(aCompany)
	//if *aCompany.Name != "" {
	if aCompany.Name != nil {
		fmt.Println("A company already exists!!!!")
		helpers.ErrorResponse(w, "A Company with name = "+*aCompany.Name+" already exists in the database", http.StatusBadRequest)
		return
	}

	if err := helpers.InsertOne(
		client,
		"companyDB",
		"companies",
		filter,
		company,
		"company"); err != nil {
		fmt.Println("Internal error ", err)
		helpers.ErrorResponse(w, "Error inserting: "+err.Error(), http.StatusInternalServerError)
		return
	}
	helpers.RespondWithJSON(w, company)

	//w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Company was created succesfully")

	// Produce an event to KAFKA
	eventMsg := "Company with name= " + *company.Name + " was created"
	if err := helpers.ProdKafkaMsg(eventMsg, kafkaConfig, kafkaProducer); err != nil {
		fmt.Println("error in kafka produce: ", err)
	}

}

/*
func checkJWT(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		//return errors.New("Missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprint(w, "Invalid token")
		return errors.New("Invalid token")
	}

	fmt.Println("Welcome to the protected area")
	return nil
}
*/
