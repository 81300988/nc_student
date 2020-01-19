package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Student struct {
	ID        int    `json:"id" bson:"id,omitempty"`
	FirstName string `json:"first_name" bson:"FirstName"`
	LastName  string `json:"last_name" bson:"LastName"`
	Age       int    `json:"age" bson:"Age"`
	ClassName string `json:"class_name" bson:"ClassName"`
	Email     string `json:"email" bson:"Email"`
}
type Error struct {
	Code int
	Msg  string
}
type UserClaims struct {
	UserID int    `json: "uid"`
	Phone  string `json: "p"`
	Email  string `json: "e"`
	jwt.StandardClaims
}
type Count struct {
	ID    int `bson:"id" json:"id"`
	Value int `bson:"value" json:"value"`
}

type Log struct {
	Date   time.Time `bson:"Date" json:"date"`
	Host   string    `bson:"Host" json:"host"`
	Path   string    `bson:"Path" json:"path"`
	Method string    `bson:"Method" json:"method"`
	Status int       `bson:"Status" json:"status"`
}
