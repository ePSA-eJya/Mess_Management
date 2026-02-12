package entities

type AdminType string

const (
	Office AdminType = "Office"
	Mess   AdminType = "Mess"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	AdminType AdminType `gorm:"type:admin_type;default:'OFFICE'" json:"admin_type"`
	Hostel    string    `gorm:"size:100;not null" json:"hostel"`
	MessNo    uint      `gorm:"not null" json:"mess_no"`
	Phone     string    `gorm:"size:15" json:"phone"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
}
