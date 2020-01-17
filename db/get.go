package database

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"nc_student.com/v1/model"
)

func GetAllStudent() ([]model.Student, error) {
	var students []model.Student
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	cur.All(ctx, &students)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor once finished
	defer cur.Close(ctx)
	return students, nil
}

func GetStudentbyID(id string) (interface{}, error) {
	var student model.Student
	collection := Client.Database("homework2").Collection("student")
	_id, _ := strconv.Atoi(id)
	filter := bson.M{"id": _id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func FindNameStartsWith(student *model.Student) ([]model.Student, error) {
	var students []model.Student
	// var filter bson.D
	var filter []bson.M
	v := reflect.ValueOf(*student)
	typeOfS := v.Type()
	fmt.Println(v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		if fieldValue := fmt.Sprintf("%v", v.Field(i).Interface()); fieldValue != "" && fieldValue != "0" && typeOfS.Field(i).Name != "ID" {
			filter = append(filter, bson.M{
				typeOfS.Field(i).Name: primitive.Regex{
					Pattern: "^" + fieldValue,
					Options: "i",
				},
			})
		}
	}
	//keyWord := student.FirstName
	collection := Client.Database("homework2").Collection("student")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err2 := collection.Find(ctx, bson.M{"$and": filter})
	if err2 != nil {
		return nil, err2
	}
	cur.All(ctx, &students)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor once finished
	defer cur.Close(ctx)
	return students, nil
}
