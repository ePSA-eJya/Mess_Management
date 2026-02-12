package entities

import (
	"time"
)

type SemesterType string

const (
	Odd  SemesterType = "ODD"
	Even SemesterType = "EVEN"
)

type Semester struct {
	SemesterID   uint         `gorm:"primaryKey" json:"semester_id"`
	AcademicYear string       `gorm:"size:100" json:"academic_year"`
	SemesterType SemesterType `gorm:"semester_type" json:"semester_type"`
	StartDate    time.Time    `gorm:"type:date;" json:"start_date"`
	EndDate      time.Time    `gorm:"type:date;" json:"end_date"`
}
