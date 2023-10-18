package models

import (
	"reflect"
	"testing"
)

func TestAccounts_Map(t *testing.T) {
	type fields struct {
		ID                    int
		StudentID             int
		TeacherID             int
		TeacherPay            float64
		StudentFee            float64
		IsDefault             bool
		OnInstallment         bool
		TotalInstallments     int
		RemainingInstallments int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - account struct to map",
			fields: fields{
				ID:                    1,
				StudentID:             1,
				TeacherID:             0,
				TeacherPay:            0.0,
				StudentFee:            3000.0,
				IsDefault:             false,
				OnInstallment:         false,
				TotalInstallments:     0,
				RemainingInstallments: 0,
			},
			want: map[string]interface{}{
				"id":                     1,
				"student_id":             1,
				"teacher_id":             0,
				"teacher_pay":            0.0,
				"student_fee":            3000.0,
				"is_default":             false,
				"on_installment":         false,
				"total_installments":     0,
				"remaining_installments": 0,
			},
		},
		{
			name: " success - account struct to map with fewer fields",
			fields: fields{
				ID:         2,
				TeacherID:  2,
				TeacherPay: 30000.0,
			},
			want: map[string]interface{}{
				"id":                     2,
				"student_id":             0,
				"teacher_id":             2,
				"teacher_pay":            30000.0,
				"student_fee":            0.0,
				"is_default":             false,
				"on_installment":         false,
				"total_installments":     0,
				"remaining_installments": 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Accounts{
				ID:                    tt.fields.ID,
				StudentID:             tt.fields.StudentID,
				TeacherID:             tt.fields.TeacherID,
				TeacherPay:            tt.fields.TeacherPay,
				StudentFee:            tt.fields.StudentFee,
				IsDefault:             tt.fields.IsDefault,
				OnInstallment:         tt.fields.OnInstallment,
				TotalInstallments:     tt.fields.TotalInstallments,
				RemainingInstallments: tt.fields.RemainingInstallments,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccounts_Names(t *testing.T) {
	type fields struct {
		ID                    int
		StudentID             int
		TeacherID             int
		TeacherPay            float64
		StudentFee            float64
		IsDefault             bool
		OnInstallment         bool
		TotalInstallments     int
		RemainingInstallments int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:                    2,
				StudentID:             2,
				TeacherID:             0,
				TeacherPay:            0.0,
				StudentFee:            3000.0,
				IsDefault:             true,
				OnInstallment:         true,
				TotalInstallments:     0,
				RemainingInstallments: 0,
			},
			want: []string{"id", "student_id", "teacher_id", "teacher_pay", "student_fee", "is_default", "on_installment", "total_installments", "remaining_installments"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Accounts{
				ID:                    tt.fields.ID,
				StudentID:             tt.fields.StudentID,
				TeacherID:             tt.fields.TeacherID,
				TeacherPay:            tt.fields.TeacherPay,
				StudentFee:            tt.fields.StudentFee,
				IsDefault:             tt.fields.IsDefault,
				OnInstallment:         tt.fields.OnInstallment,
				TotalInstallments:     tt.fields.TotalInstallments,
				RemainingInstallments: tt.fields.RemainingInstallments,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
