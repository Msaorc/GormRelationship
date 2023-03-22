package database

import (
	"github.com/Msaorc/GoRelationship/entity"
	"gorm.io/gorm"
)

func CreateExpenseLevel(db *gorm.DB, description string) (*entity.ExpenseLevel, error) {
	var expenseLevel = entity.ExpenseLevel{Description: description}
	err := db.Create(&expenseLevel).Error
	return &expenseLevel, err
}
