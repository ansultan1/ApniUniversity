package models

import "github.com/fatih/structs"

type Teachers struct {
	ID         int    `json:"id" bson:"_id" structs:"id" db:"id"`
	Name       string `json:"name" bson:"name" structs:"name" db:"name"`
	Department string `json:"department" bson:"department" structs:"department" db:"department"`
}

// Map converts structs to a map representation
func (t *Teachers) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of Teachers model
func (t *Teachers) Names() []string {
	fields := structs.Fields(t)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
