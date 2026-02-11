package entities

type StudentStatus string

const (
	Active   StudentStatus = "ACTIVE"
	Inactive StudentStatus = "INACTIVE"
)

type Student struct {
	Roll   uint          `gorm:"primaryKey" json:"roll"`
	Name   string        `gorm:"size:100;not null" json:"name"`
	Hostel string        `gorm:"size:100;not null" json:"hostel"`
	RoomNo uint          `gorm:"not null" json:"room_no"`
	MessNo uint          `gorm:"not null" json:"mess_no"`
	Phone  string        `gorm:"size:15" json:"phone"`
	Email  string        `gorm:"size:255;unique;not null" json:"email"`
	Status StudentStatus `gorm:"type:student_status;default:'ACTIVE'" json:"status"`
}
