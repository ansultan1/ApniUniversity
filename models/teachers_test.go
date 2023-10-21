package models

import (
	"reflect"
	"testing"
)

func TestTeachers_Map(t1 *testing.T) {
	type fields struct {
		ID         int
		Name       string
		Department string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - teacher struct to map",
			fields: fields{
				ID:         1,
				Name:       "ahmad",
				Department: "IT",
			},
			want: map[string]interface{}{
				"id":         1,
				"name":       "ahmad",
				"department": "IT",
			},
		},
		{
			name: " success - teacher struct to map with fewer fields",
			fields: fields{
				ID:         0,
				Department: "IT",
			},
			want: map[string]interface{}{
				"id":         0,
				"name":       "",
				"department": "IT",
			},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Teachers{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Department: tt.fields.Department,
			}
			if got := t.Map(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeachers_Names(t1 *testing.T) {
	type fields struct {
		ID         int
		Name       string
		Department string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - return list of teacher fields",
			fields: fields{
				ID:         2,
				Name:       "khan",
				Department: "CS",
			},
			want: []string{"id", "name", "department"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Teachers{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Department: tt.fields.Department,
			}
			if got := t.Names(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
