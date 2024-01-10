package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckName(collection *mongo.Collection, name string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "name", Value: name},
	}

	var result bson.D

	err := collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		fmt.Println("Error in finding the name you want :", err)
	}

	if result == nil {
		fmt.Println("Name not found")
		fmt.Println("")
		return false
	}

	return true
}
