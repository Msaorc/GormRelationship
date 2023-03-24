package entity

import "github.com/Msaorc/GoRelationship/pkg/entity"

type ExpenseOrigin struct {
	ID          entity.ID `gorm:"primaryKey"`
	Description string
}
