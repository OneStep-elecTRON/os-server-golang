package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password []byte `json:"-"`
}
