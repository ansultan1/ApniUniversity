package models

import (
	"reflect"
	"testing"
	"time"
)

func TestAccount_Map(t *testing.T) {
	type fields struct {
		ID                 int
		Type               AccountType
		TeacherPay         float64
		StudentFee         float64
		IsDefault          bool
		OnInstallment      bool
		InstallmentDetails Installment
		CreatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - convert Account struct to map",
			fields: fields{
				ID:                 1,
				Type:               "student",
				TeacherPay:         0.0,
				StudentFee:         2000.0,
				IsDefault:          false,
				OnInstallment:      false,
				InstallmentDetails: Installment{0, 0},
				CreatedAt:          time.Time{},
			},
			want: map[string]interface{}{
				"id":                 1,
				"type":               "student",
				"teacher_pay":        0.0,
				"student_fee":        2000.0,
				"is_default":         false,
				"on_installment":     false,
				"installmentDetails": Installment{0, 0},
				"created_at":         time.Time{},
			},
		},
		{
			name: "success - convert  account struct to map2",
			fields: fields{
				ID:                 2,
				Type:               "student",
				TeacherPay:         0.0,
				StudentFee:         2000.0,
				IsDefault:          false,
				OnInstallment:      false,
				InstallmentDetails: Installment{0, 0},
				CreatedAt:          time.Time{},
			},
			want: map[string]interface{}{
				"id":                  2,
				"type":                "student",
				"teacher_pay":         0.0,
				"student_fee":         2000.0,
				"is_default":          false,
				"on_installment":      false,
				"installment_details": Installment{0, 0},
				"created_at":          time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Accounts{
				ID:                 tt.fields.ID,
				Type:               tt.fields.Type,
				TeacherPay:         tt.fields.TeacherPay,
				StudentFee:         tt.fields.StudentFee,
				IsDefault:          tt.fields.IsDefault,
				OnInstallment:      tt.fields.OnInstallment,
				InstallmentDetails: tt.fields.InstallmentDetails,
				CreatedAt:          tt.fields.CreatedAt,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Names(t *testing.T) {
	type fields struct {
		ID                 int
		Type               AccountType
		TeacherPay         float64
		StudentFee         float64
		IsDefault          bool
		OnInstallment      bool
		InstallmentDetails Installment
		CreatedAt          time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of account struct fields",
			fields: fields{
				ID:                 1,
				Type:               "student",
				TeacherPay:         0,
				StudentFee:         1500.0,
				IsDefault:          false,
				OnInstallment:      false,
				InstallmentDetails: Installment{0, 0},
				CreatedAt:          time.Time{},
			},
			want: []string{"id", "type", "teacher_pay", "student_fee", "is_default", "on_installment", "installment_details", "created_at"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Accounts{
				ID:                 tt.fields.ID,
				Type:               AccountType(tt.fields.Type),
				TeacherPay:         tt.fields.TeacherPay,
				StudentFee:         tt.fields.StudentFee,
				IsDefault:          tt.fields.IsDefault,
				OnInstallment:      tt.fields.OnInstallment,
				InstallmentDetails: Installment{0, 0},
				CreatedAt:          tt.fields.CreatedAt,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
