package database

import (
	"github.com/Msaorc/GoRelationship/entity"
	pkgID "github.com/Msaorc/GoRelationship/pkg/entity"
	"gorm.io/gorm"
)

func CreateExpense(db *gorm.DB, description string, value float64, note string, expenseLevelID pkgID.ID, expenseOriginID pkgID.ID) error {
	var expense = entity.Expense{
		ID:          pkgID.NewID(),
		Description: description,
		Value:       value, Note: note,
		ExpenseLevelID:  expenseLevelID,
		ExpenseOriginID: expenseOriginID,
	}
	return db.Create(&expense).Error
}
