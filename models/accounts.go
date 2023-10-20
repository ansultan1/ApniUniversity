package models

import (
	"github.com/fatih/structs"
)

type Accounts struct {
	ID                    int     `json:"id" bson:"_id" structs:"id" db:"id"`
	Type                  string  `json:"type" bson:"type" structs:"type" db:"type"`
	StudentFee            float64 `json:"studentFee" bson:"student_fee" structs:"student_fee" db:"student_fee"`
	TeacherPay            float64 `bson:"teacherPay" bson:"teacher_pay" structs:"teacher_pay" db:"teacher_pay"`
	IsDefault             bool    `json:"isDefault" bson:"is_default" structs:"is_default" db:"is_default"`
	OnInstallment         bool    `json:"onInstallment" bson:"on_installment" structs:"on_installment" db:"on_installment"`
	TotalInstallments     int     `json:"totalInstallments" bson:"total_installments" structs:"total_installments" db:"total_installments"`
	RemainingInstallments int     `json:"remainingInstallments" bson:"remaining_installments" structs:"remaining_installments" db:"remaining_installments"`
}

func (a *Accounts) Map() map[string]interface{} {
	return structs.Map(a)
}
