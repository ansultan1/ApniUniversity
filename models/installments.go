package models

type Installment struct {
	TotalInstallments     int `json:"totalInstallments" bson:"total_installments" db:"total_installments"`
	RemainingInstallments int `json:"remainingInstallments" bson:"remaining_installments" db:"remaining_installments"`
}
