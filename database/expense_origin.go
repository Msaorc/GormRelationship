package database

import (
	"github.com/Msaorc/GoRelationship/entity"
	"gorm.io/gorm"
)

func CreateExpenseOrigin(db *gorm.DB, description string) (*entity.ExpenseOrigin, error) {
	var expenseOrigin = entity.ExpenseOrigin{Description: description}
	return &expenseOrigin, db.Create(&expenseOrigin).Error
}
