package models

import "github.com/fatih/structs"

type Students struct {
	ID         int    `json:"ID" bson:"_id" structs:"id" db:"id"`
	Name       string `json:"name" bson:"name" structs:"name" db:"name"`
	FatherName string `json:"fatherName" bson:"father_name" structs:"father_name"`
}

// Map converts structs to a map representation
func (s *Students) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of students model
func (s *Students) Names() []string {
	fields := structs.Fields(s)
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
