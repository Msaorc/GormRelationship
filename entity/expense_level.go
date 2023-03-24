package entity

import "github.com/Msaorc/GoRelationship/pkg/entity"

type ExpenseLevel struct {
	ID          entity.ID `gorm:"primaryKey"`
	Description string
}
