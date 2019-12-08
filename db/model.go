package database

type Student struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	FirstName string      `json:"first_name" bson:"FirstName"`
	LastName  string      `json:"last_name" bson:"LastName"`
	Age       int         `json:"age" bson:"Age"`
	ClassName string      `json:"class_name" bson:"ClassName"`
	Email     string      `json:"email" bson:"Email"`
}
type Error struct {
	Code int
	Msg  string
}
