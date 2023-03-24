package main

import (
	"fmt"

	"github.com/Msaorc/GoRelationship/database"
	"github.com/Msaorc/GoRelationship/entity"
)

func main() {
	db := database.GetDatabase()
	db.AutoMigrate(&entity.ExpenseLevel{}, &entity.ExpenseOrigin{}, &entity.Expense{})
	expenseLevel, err := database.CreateExpenseLevel(db, "Low")
	if err != nil {
		panic(err)
	}
	expenseOrigin, err := database.CreateExpenseOrigin(db, "Pao de Açucar")
	if err != nil {
		panic(err)
	}
	err = database.CreateExpense(db, "Copo", 120.00, "Pocoto de Asa", expenseLevel.ID, expenseOrigin.ID)
	if err != nil {
		panic(err)
	}
	var expenses []entity.Expense
	db.Preload("ExpenseLevel").Preload("ExpenseOrigin").Find(&expenses)
	for _, expense := range expenses {
		fmt.Println(expense.ID, expense.Description, expense.Value, expense.ExpenseLevel.Description, expense.ExpenseOrigin.Description)
	}
}
