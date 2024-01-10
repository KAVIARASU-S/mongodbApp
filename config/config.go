package config

import (
	"appplay/constants"
	"context"
	"fmt"
	

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func ConnectDatabase(ctx context.Context, Mongo *mongo.Client) (*mongo.Client){
	Mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(constants.ConnectionString))
	if err != nil {
		fmt.Println("Error connecting to mongoDB :", err)
		return nil
	} else {
		fmt.Println("MongoDB connected succesfully")
	}

	return Mongo
}
