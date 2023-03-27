package entity

import (
	"github.com/Msaorc/GoRelationship/pkg/entity"
	"gorm.io/gorm"
)

type Expense struct {
	ID              entity.ID `gorm:"primaryKey"`
	Description     string
	Value           float64
	Note            string
	ExpenseLevelID  entity.ID `gorm:"size:36"`
	ExpenseLevel    ExpenseLevel
	ExpenseOriginID entity.ID `gorm:"size:36"`
	ExpenseOrigin   ExpenseOrigin
	gorm.Model
}
