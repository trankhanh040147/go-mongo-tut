package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/trankhanh040147/go-mongo-tut/modules/restaurant/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// start: load env & connect to db
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// end: load env & connect to db

	// test query
	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"

	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	findQuery02(client)
}

// find one
func findQuery01(client *mongo.Client) {
	var result model.Restaurant
	coll := client.Database("sample_restaurants").Collection(result.CollName())

	// Creates a query filter to match documents in which the "name" is
	// "Bagels N Buns"
	filter := bson.D{{"name", "Bagels N Buns"}}

	// Retrieves the first matching document
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}

	fmt.Println(result)
}

// find many
func findQuery02(client *mongo.Client) {
	coll := client.Database("sample_restaurants").Collection("restaurants")

	// Creates a query filter to match documents in which the "cuisine"
	// is "Italian"
	filter := bson.D{{"cuisine", "Italian"}}

	// Retrieves documents that match the query filer
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []model.Restaurant
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Prints the results of the find operation as structs
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}
