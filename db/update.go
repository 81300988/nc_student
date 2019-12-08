package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func UpdateOneStudent(student *Student) (interface{}, error) {
	collection := Client.Database("homework2").Collection("student")
	id, err := primitive.ObjectIDFromHex(student.ID.(string))
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", id)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Email": student.Email}}
	updateResult, err1 := collection.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		return nil, err1
	}
	return updateResult, nil
}
