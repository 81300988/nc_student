package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DeleteOneStudent(student *Student) (interface{}, error) {
	collection := Client.Database("homework2").Collection("student")
	id, err := primitive.ObjectIDFromHex(student.ID.(string))
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", id)
	}
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult, nil
}
