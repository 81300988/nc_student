package database

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"
	"nc_student.com/v1/model"
)

func DeleteOneStudent(student *model.Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	// id, err := primitive.ObjectIDFromHex(student.ID.(string))
	// if err != nil {
	// 	fmt.Println("ObjectIDFromHex ERROR", err)
	// } else {
	// 	fmt.Println("ObjectIDFromHex:", id)
	// }
	filter := bson.M{"id": student.ID}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult, nil
}
