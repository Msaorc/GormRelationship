package database

import (
	"github.com/Msaorc/GoRelationship/entity"
	pkgID "github.com/Msaorc/GoRelationship/pkg/entity"
	"gorm.io/gorm"
)

func CreateExpenseLevel(db *gorm.DB, description string) (*entity.ExpenseLevel, error) {
	var expenseLevel = entity.ExpenseLevel{
		ID:          pkgID.NewID(),
		Description: description}
	err := db.Create(&expenseLevel).Error
	return &expenseLevel, err
}
