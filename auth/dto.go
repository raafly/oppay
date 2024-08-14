package auth

import "time"

type UserReq struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Pin             int       `json:"key"`
	Email           string    `json:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
}
