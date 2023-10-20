package models

import (
	"github.com/fatih/structs"
	"time"
)

type Accounts struct {
	ID                   int         `json:"id" bson:"_id" structs:"id" db:"id"`
	Type                 AccountType `json:"type" bson:"type" structs:"type" db:"type"`
	TeacherPay           float64     `json:"teacherPay" bson:"teacher_pay" structs:"teacher_pay" db:"teacher_pay"`
	StudentFee           float64     `json:"studentFee" bson:"student_fee" structs:"student_fee" db:"student_fee"`
	IsDefault            bool        `json:"isDefault" bson:"is_default" structs:"is_default" db:"is_default"`
	StudentOnInstallment bool        `json:"studentOnInstallment" bson:"student_on_installment" structs:"student_on_installment"`
	InstallmentDetails   Installment `json:"installmentDetails" bson:"installment_details" structs:"installment_details" db:"installment_details"`
	CreatedAt            time.Time   `json:"createdAt" bson:"created_at" structs:"created_at" db:"created_at"`
}

func (a *Accounts) Map() map[string]interface{} {
	return structs.Map(a)
}

// Names returns the field names of Accounts model
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
