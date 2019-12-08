package database

import (
	"context"
)

func AddOneStudent(student *Student) (interface{}, error) {
	collection := Client.Database("homework2").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), &student)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
