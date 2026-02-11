package entities

import "time"

type MealType string

const (
	Breakfast MealType = "BREAKFAST"
	Lunch     MealType = "LUNCH"
	Dinner    MealType = "DINNER"
)

type MealCancellationRecord struct {
	Roll       uint      `gorm:"" json:"roll"`
	SemesterID uint      `gorm:"" json:"semester_id"`
	MealType   MealType  `gorm:"type:meal_type;default:'BREAKFAST'" json:"meal_type"`
	Date       time.Time `gorm:"" json:"date"`
}
