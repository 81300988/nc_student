package database

import (
	"context"
	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func GetAllStudent() ([]Student, error) {
	var students []Student
	collection := Client.Database("homework2").Collection("student")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Student
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		students = append(students, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())
	return students, nil
}

func GetOneStudent(id string) (interface{}, error) {
	var student Student
	collection := Client.Database("homework2").Collection("student")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", _id)
	}
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}
func FindNameStartsWith(student *Student) ([]Student, error) {
	var students []Student
	keyWord := "^" + student.FirstName
	collection := Client.Database("homework2").Collection("student")
	cur, err2 := collection.Find(context.TODO(), bson.M{
		"FirstName": primitive.Regex{
			Pattern: keyWord,
			Options: "i",
		},
	})
	if err2 != nil {
		return nil, err2
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Student
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		students = append(students, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())
	return students, nil
}
