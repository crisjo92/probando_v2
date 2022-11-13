package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOption = options.Client().ApplyURI("mongodb+srv://dbCrist:<password>@probandomongo.h1ugsfn.mongodb.net/?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Conexion Existosa con la BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
