package database

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
	"nc_student.com/v1/model"
)

func UpdateOneStudent(student *model.Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	// id, err := primitive.ObjectIDFromHex(student.ID.(string))
	// if err != nil {
	// 	fmt.Println("ObjectIDFromHex ERROR", err)
	// } else {
	// 	fmt.Println("ObjectIDFromHex:", id)
	// }
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": student.ID}
	update := bson.M{"$set": bson.M{"Email": student.Email}}
	updateResult, err1 := collection.UpdateOne(ctx, filter, update)
	if err1 != nil {
		return nil, err1
	}
	return updateResult, nil
}
