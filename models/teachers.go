package models

import "github.com/fatih/structs"

type Teachers struct {
	ID         int    `json:"id" bson:"_id" structs:"id" db:"id"`
	AccountID  int    `json:"accountId" bson:"account_id" structs:"account_id" db:"account_id"`
	Name       string `json:"name" bson:"name" structs:"name" db:"name"`
	Department string `json:"department" bson:"department" structs:"department" db:"department"`
}

func (t *Teachers) Map() map[string]interface{} {
	return structs.Map(t)
}
