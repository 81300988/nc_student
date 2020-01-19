package database

import (
	"context"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
	"nc_student.com/v1/model"
)

func AddOneStudent(student *model.Student) (interface{}, error) {
	var count []model.Count
	collection := Client.Database(DbName).Collection(ColName)
	collectionCount := Client.Database(DbName).Collection("count")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collectionCount.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	cur.All(ctx, &count)
	student.ID = count[0].Value + 1
	filter := bson.M{"id": count[0].ID}
	update := bson.M{"$set": bson.M{"value": student.ID}}
	_, err1 := collectionCount.UpdateOne(ctx, filter, update)
	insertResult, err := collection.InsertOne(ctx, &student)
	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}

	return insertResult, nil

}

func AddLog(log *model.Log) (interface{}, error) {
	collection := Client.Database(DbName).Collection("log")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	insertResult, err := collection.InsertOne(ctx, log)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
