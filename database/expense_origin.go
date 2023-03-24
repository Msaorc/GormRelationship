package database

import (
	"github.com/Msaorc/GoRelationship/entity"
	pkgID "github.com/Msaorc/GoRelationship/pkg/entity"
	"gorm.io/gorm"
)

func CreateExpenseOrigin(db *gorm.DB, description string) (*entity.ExpenseOrigin, error) {
	var expenseOrigin = entity.ExpenseOrigin{
		ID: pkgID.NewID(),
		Description: description}
	return &expenseOrigin, db.Create(&expenseOrigin).Error
}
