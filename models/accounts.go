package models

import "github.com/fatih/structs"

type Accounts struct {
	ID                    int     `json:"id" bson:"_id" structs:"id" db:"id"`
	StudentID             int     `json:"studentID" bson:"student_id" structs:"student_id" db:"student_id"`
	TeacherID             int     `json:"teacherID" bson:"teacher_id" structs:"teacher_id" db:"teacher_id"`
	TeacherPay            float64 `bson:"teacherPay" bson:"teacher_pay" structs:"teacher_pay" db:"teacher_pay"`
	StudentFee            float64 `json:"studentFee" bson:"student_fee" structs:"student_fee" db:"student_fee"`
	IsDefault             bool    `json:"isDefault" bson:"is_default" structs:"is_default" db:"is_default"`
	OnInstallment         bool    `json:"onInstallment" bson:"on_installment" structs:"on_installment" db:"on_installment"`
	TotalInstallments     int     `json:"totalInstallments" bson:"total_installments" structs:"total_installments" db:"total_installments"`
	RemainingInstallments int     `json:"remainingInstallments" bson:"remaining_installments" structs:"remaining_installments" db:"remaining_installments"`
}

// Map converts structs to a map representation
func (a *Accounts) Map() map[string]interface{} {
	return structs.Map(a)
}

// Names returns the field names of Accounts model
func (a *Accounts) Names() []string {
	fields := structs.Fields(a)
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
