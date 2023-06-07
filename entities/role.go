package entities

type Role struct {
	ID       uint   `gorm:"primary_key"`
	RoleName string `gorm:"column:role_name"`
}
