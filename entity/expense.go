package entity

import (
	"github.com/Msaorc/GoRelationship/pkg/entity"
)

type Expense struct {
	ID             entity.ID `gorm:"primaryKey"`
	Description    string
	Value          float64
	Note           string
	ExpenseLevelID entity.ID
	ExpenseLevel
	ExpenseOriginID entity.ID
	ExpenseOrigin
}
