package entities

import (
	"github.com/google/uuid"
)

type MonthlyBill struct {
	BillID         uuid.UUID `gorm:"primaryKey" json:"bill_id"`
	Roll           uint      `gorm:"" json:"roll"`
	Month          string    `gorm:"size:100" json:"month"`
	SemesterID     uint      `gorm:"" json:"semester_id"`
	BreakfastCount uint      `gorm:"" json:"breakfast_count"`
	LunchCount     uint      `gorm:"" json:"lunch_count"`
	DinnerCount    uint      `gorm:"" json:"dinner_count"`
	TotalBill      float64   `gorm:"type:decimal(10,2);" json:"total_bill"`
}
