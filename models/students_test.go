package models

import (
	"reflect"
	"testing"
)

func TestStudents_Map(t *testing.T) {
	type fields struct {
		ID         int
		Name       string
		FatherName string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - student struct to map",
			fields: fields{
				ID:         1,
				Name:       "ahmed",
				FatherName: "khan",
			},
			want: map[string]interface{}{
				"id":          1,
				"name":        "ahmed",
				"father_name": "khan",
			},
		},
		{
			name: " success - student struct to map with fewer fields",
			fields: fields{
				Name:       "ahmed",
				FatherName: "khan",
			},
			want: map[string]interface{}{
				"id":          0,
				"name":        "ahmed",
				"father_name": "khan",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Students{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				FatherName: tt.fields.FatherName,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_Names(t *testing.T) {
	type fields struct {
		ID         int
		Name       string
		FatherName string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:         2,
				Name:       "ahmed",
				FatherName: "khan",
			},
			want: []string{"id", "name", "father_name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Students{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				FatherName: tt.fields.FatherName,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
