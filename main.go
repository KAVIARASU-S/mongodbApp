package main

import (
	"appplay/config"
	"appplay/constants"
	"appplay/menu"
	"appplay/services"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
    Mongo *mongo.Client
    collection *mongo.Collection
)



func GetCollection(Mongo *mongo.Client)(*mongo.Collection){
    fmt.Println("inside collection")
    collection= Mongo.Database(constants.Database).Collection(constants.Collection)
    return collection
}


func main() {
    ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
    defer cancel()

    fmt.Println("Welcome to MongoDB service")
    Mongo=config.ConnectDatabase(ctx, Mongo)

    fmt.Println("going to collection")
    collection=GetCollection(Mongo)

	menu.Menu()

    var input int

    for {
        fmt.Println("Enter your choice")
        fmt.Scanln(&input)

        if input==7{
            fmt.Println("Closed successfully")
            break
        }

        switch input {
        case 1: services.DisplayName(collection)
        case 2: services.InsertName(collection)
        case 3: services.FindName(collection)
        case 4: services.UpdateName(collection)
        case 5: services.DeleteName(collection)
        case 6: menu.Menu()
        default: fmt.Println("Enter correct input")
        }
        
    
    }
}
