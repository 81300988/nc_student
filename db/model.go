package database

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassName string `json:"class_name"`
	Email     string `json:"email"`
}
type Error struct {
	Code int
	Msg  string
}
