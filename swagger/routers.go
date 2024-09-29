package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fomichalopoulos/companiesMicroService/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client
var kafkaConfig models.KafkaConfig
var kafkaProducer *kafka.Producer

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Array of Route structs
type Routes []Route

// NewRouter returns a new router
func NewRouter(client *mongo.Client, kafkacfg models.KafkaConfig, producer *kafka.Producer) *mux.Router {
	//	envcfg, err := configs.InitConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println("created mongoClient in router: ", client)
	//	MongoClient = client

	MongoClient = client
	kafkaConfig = kafkacfg
	kafkaProducer = producer

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	/*
		MongoClient = client
		kafkaConfig = kafkacfg
		kafkaProducer = producer
	*/
	fmt.Println("client = ", MongoClient)

	return router
}

var routes = Routes{

	Route{
		"Login",
		strings.ToUpper("POST"),
		"/company/userLogin",
		Login,
	},
	Route{
		"CreateCompany",
		strings.ToUpper("Post"),
		"/company/create",
		CreateCompany,
	},
	Route{
		"GetCompany",
		strings.ToUpper("Get"),
		"/company",
		GetCompany,
	},
	Route{
		"DelCompany",
		strings.ToUpper("Delete"),
		"/company",
		DelCompany,
	},
	Route{
		"PatchCompany",
		strings.ToUpper("Patch"),
		"/company",
		PatchCompany,
	},
}
