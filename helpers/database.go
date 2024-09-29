package helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(
	//	w http.ResponseWriter,
	client *mongo.Client,
	databaseName string,
	collection string,
	filter primitive.D,
	insertElement interface{},
	item string,
	filterOptions ...*options.FindOneOptions) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("db = ", databaseName)
	fmt.Println("dbCollection = ", collection)

	db := client.Database(databaseName)
	dbCollection := db.Collection(collection + "Collection")

	if _, err := dbCollection.InsertOne(ctx, insertElement); err != nil {
		//ErrorResponse(w, "Error inserting kpi: "+err.Error(), http.StatusInternalServerError)
		//return nil
		return errors.New("Error inserting company: " + err.Error())
	}

	return nil

}

func UpdateUsingFilter(
	client *mongo.Client,
	//id string,
	filter primitive.M,
	databaseName string,
	updateSet primitive.M,
	collection string,
) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := client.Database(databaseName)
	dbCollection := db.Collection(collection + "Collection")

	fmt.Println("Before invoking updateOne")

	result, err := dbCollection.UpdateOne(ctx, filter, updateSet)
	if err != nil {
		//ErrorResponse(w, "Could not insert the updated "+item+" to database: "+err.Error(), http.StatusInternalServerError)
		return err
	}
	if result.ModifiedCount == 0 {
		//ErrorResponse(w, "No "+item+" was updated", http.StatusNotModified)
		return errors.New("No company was updated")
	}

	return nil

}

func DeleteOne(
	client *mongo.Client,
	databaseName string,
	collection string,
	filter primitive.D,
	item string,
	filterOptions ...*options.FindOneOptions) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := client.Database(databaseName)
	dbCollection := db.Collection(collection + "Collection")

	deleteResut, err := dbCollection.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Println("Error in deleting one: ", err)
		return errors.New("Error deleting: " + item + " " + err.Error())
	}

	if deleteResut.DeletedCount == 0 {
		return errors.New("No company exists on the database")
	}

	return nil

}

func CheckIfExists(
	client *mongo.Client,
	databaseName string,
	collection string,
	filter primitive.D,
	result interface{},
	item string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("db = ", databaseName)
	fmt.Println("dbCollection = ", collection)

	db := client.Database(databaseName)
	dbCollection := db.Collection(collection + "Collection")

	if err := dbCollection.FindOne(ctx, filter).Decode(result); err != nil && err != mongo.ErrNoDocuments {
		//helpers.ErrorResponse(w, "Checking "+item+" database error: "+err.Error(), http.StatusBadRequest)

		return errors.New("Checking " + item + " database error: " + err.Error())
	}
	return nil

}

func FetchUsingProperty(
	client *mongo.Client,
	databaseName string,
	collection string,
	filter primitive.D,
	result interface{},
	item string,
	filterOptions ...*options.FindOneOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("db = ", databaseName)
	fmt.Println("dbCollection = ", collection)

	db := client.Database(databaseName)
	dbCollection := db.Collection(collection + "Collection")
	if err := dbCollection.FindOne(ctx, filter, filterOptions...).Decode(result); err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		} else {
			return err
		}
	}

	return nil

}
