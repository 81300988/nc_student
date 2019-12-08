package database

import (
	"context"
	"log"

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
	cur.Close(context.TODO())
	return students, nil
}
