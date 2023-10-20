package models

type Students struct {
	ID         int    `json:"ID" bson:"id" structs:"id" db:"id""`
	AccountID  int    `json:"id" bson:"_id" structs:"id" db:"id"`
	Name       string `json:"name" bson:"name" structs:"name" db:"name"`
	FatherName string `json:"fatherName" bson:"father_name" structs:"father_name"`
}
