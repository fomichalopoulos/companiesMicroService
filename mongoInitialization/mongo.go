package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Mongo_URI string `env:"MONGO_COMPANY_URI"`
}

/*
func loadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Sprintln("no .env file present")

	}

	return nil

}
*/

func main() {
	mongocfg := MongoConfig{}
	if err := env.Parse(&mongocfg); err != nil {
		log.Fatalf("Error in parsing : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongocfg.Mongo_URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Error in instantiating mongo client %v", err)
	}
	companyDB := client.Database("companyDB")
	companiesCollection := companyDB.Collection("companiesCollection")

	companies := //interface{}{
		bson.D{
			{Key: "name", Value: "XM"},
			{Key: "description", Value: "This is a trading company"},
			{Key: "amount", Value: 3},
			{Key: "registered", Value: true},
			{Key: "type", Value: "Corporations"},
		}
	//}

	res, err := companiesCollection.InsertOne(ctx, companies)
	if err != nil {
		log.Fatalf("error when inserting companies in companies collection %s", err)
	}

	fmt.Println("company added succesfully in the companies collection ", res)
}
