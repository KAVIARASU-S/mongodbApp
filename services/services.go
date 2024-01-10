package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DisplayName(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		fmt.Println("")
		fmt.Println("Error displaying the names :", err)
	}

	var NamesList []bson.D

	cursor.All(ctx, &NamesList)

	for _, i := range NamesList {
		fmt.Println(i)
	}

}

func InsertName(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	fmt.Println("Enter the name to insert")

	var name string
	fmt.Scanln(&name)

	collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: name},
	})

	fmt.Println("Inserted Successfully")
}

func FindName(collection *mongo.Collection) {
	fmt.Println("Enter the name you want to find")

	var name string
	fmt.Scanln(&name)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	filter := bson.D{
		{Key: "name", Value: name},
	}

	var result bson.D

	err:=collection.FindOne(ctx, filter).Decode(&result)

	if err!=nil{
		fmt.Println("Error in finding the name :",err)
	}

	if result == nil {
		fmt.Println("Name not found")
		fmt.Println("")
	} else {
		fmt.Println(result)
	}
}

func UpdateName(collection *mongo.Collection) {
	var name string

	fmt.Println("")
	fmt.Println("Enter the name you want to update")
	fmt.Scanln(&name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	//Function to find wheter the entered name exists
	FoundName := CheckName(collection,name)

	if !FoundName{
		return
	}

	filter := bson.D{
		{Key: "name", Value: name},
	}

	fmt.Println("Enter the new name")
	fmt.Scanln(&name)

	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: name},
			},
		},
	}

	_,err:=collection.UpdateOne(ctx,filter,update)
	if err!=nil{
		fmt.Println("Error in updating the name :",err)
	}

}

func DeleteName (collection *mongo.Collection){
	fmt.Println("")
	fmt.Println("Enter the name you want to delete")

	var name string
	fmt.Scanln(&name)

	FoundName:=CheckName(collection,name)

	if !FoundName{
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	filter:= bson.D{
		{Key: "name",Value: name},
	}

	_,err:=collection.DeleteOne(ctx,filter)

	if err != nil {
		fmt.Println("Error deleting the name :",err)
	}else{
		fmt.Println("")
		fmt.Println("Name deleted successfully")
	}
}