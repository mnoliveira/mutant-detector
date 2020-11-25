package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mutant-detector/config"
	"mutant-detector/model"
)

const humanCollection = "humans"
const mutantCollection = "mutants"

var ctx = context.TODO()

func SaveMutant(dna model.DNADB) error {
	return saveDNA(dna, mutantCollection)
}

func SaveHuman(dna model.DNADB) error {
	return saveDNA(dna, humanCollection)
}

func saveDNA(dna model.DNADB, collectionName string) error {

	db, err := getDB()
	if err != nil {
		return err
	}

	collection := db.Collection(collectionName)

	filter := bson. M{"dna": bson. M{"$eq":dna.DNA}}

	var dnaDB model.DNADB
	err = collection.FindOne(ctx, filter).Decode(&dnaDB)
	if err != nil && err == mongo.ErrNoDocuments {
		_, err = collection.InsertOne(ctx, dna)
	}

	return err
}

func GetMutantCount() (int64, error) {
	return getDocumentCount(mutantCollection)
}

func GetHumanCount() (int64, error) {
	return getDocumentCount(humanCollection)
}

func getDocumentCount(collectionName string) (int64, error) {

	db, err := getDB()
	if err != nil {
		return 0, err
	}

	collectionMutant := db.Collection(collectionName)

	return collectionMutant.CountDocuments(ctx, bson.D{})
}

func getDB() (*mongo.Database, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	database, _ := config.Config.String("database.mongodb.name")

	return client.Database(database), nil
}

func getClient() (*mongo.Client, error) {

	host, _ := config.Config.String("database.mongodb.host")
	port, _ := config.Config.Int("database.mongodb.port")

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
