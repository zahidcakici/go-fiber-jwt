package model

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password_hash,omitempty"`
}
