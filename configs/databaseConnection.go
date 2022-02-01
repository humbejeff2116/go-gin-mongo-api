





package configs

import (
	"context"
	"fmt"
	// "log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)





// MONGODB client instance
var MongodbClient *mongo.Client = ConnectToMongDb();


func ConnectToMongDb() *mongo.Client{
	serverConfigs := SetServerConfigurations();
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel()
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverConfigs.DatabaseURI));
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
	// Ping the primary
	err = client.Ping(ctx, readpref.Primary());
	if  err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged mongo database");
	return client;
}

//used to get a database  collection
func GetCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
    collection := client.Database(databaseName).Collection(collectionName);
    return collection;
}