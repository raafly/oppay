package auth

import "time"

type User struct {
	ID                string `gorm:"primaryKey"`
	Name              string
	Email             string
	Pin               int
	VerificationToken string
	RememberToken     string
	EmailVerifiedAt   time.Time
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
}
