package main

import (
	"github.com/Msaorc/GoRelationship/database"
)

func main() {
	db := database.GetDatabase()
	expenseLevel, err := database.CreateExpenseLevel(db, "Hard")
	if err != nil {
		panic(err)
	}
	expenseOrigin, err := database.CreateExpenseOrigin(db, "Nubank")
	if err != nil {
		panic(err)
	}
	err = database.CreateExpense(db, "Corda Violão", 120.00, "Corda para violao canario aço", expenseLevel.ID, expenseOrigin.ID)
	if err != nil {
		panic(err)
	}
}
