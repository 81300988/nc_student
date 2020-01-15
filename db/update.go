package database

import (
	"context"

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
	filter := bson.M{"id": student.ID}
	update := bson.M{"$set": bson.M{"Email": student.Email}}
	updateResult, err1 := collection.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		return nil, err1
	}
	return updateResult, nil
}
