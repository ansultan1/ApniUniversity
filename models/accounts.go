package models

import (
	"github.com/fatih/structs"
	"time"
)

type Accounts struct {
	ID                    int       `json:"id" bson:"_id" structs:"id" db:"id"`
	Type                  string    `json:"type" bson:"type" structs:"type" db:"type"`
	TeacherPay            float64   `json:"teacher_pay" bson:"teacher_pay" structs:"teacher_pay" db:"teacher_pay"`
	StudentFee            float64   `json:"studentFee" bson:"student_fee" structs:"student_fee" db:"student_fee"`
	DefaultStudent        bool      `json:"defaultStudent" bson:"default_student" structs:"default_student" db:"default_student"`
	StudentOnInstallment  bool      `json:"studentsOnInstallment" bson:"student_on_installment" structs:"student_on_installment" db:"student_on_installment"`
	TotalInstallments     int       `json:"totalInstallments" bson:"total_installments" structs:"total_installments" db:"total_installments"`
	RemainingInstallments int       `json:"remainingInstallments" bson:"remaining_installments" structs:"remaining_installments" db:"remaining_installments"`
	CreatedAt             time.Time `json:"createdAt" bson:"created_at" structs:"created_at" db:"created_at"`
}

func (a *Accounts) Map() map[string]interface{} {
	return structs.Map(a)
}

// Names returns the field names of Teacher model
func (a Accounts) Names() []string {
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
