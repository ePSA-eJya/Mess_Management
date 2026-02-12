package entities

import (
	"github.com/google/uuid"
)

type SemesterBill struct {
	BillID     uuid.UUID `gorm:"primaryKey" json:"bill_id"`
	Roll       uint      `gorm:"" json:"roll"`
	SemesterID uint      `gorm:"" json:"semester_id"`
	TotalBill  float64   `gorm:"type:decimal(10,2);" json:"total_bill"`
}
