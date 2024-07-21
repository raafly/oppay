package core

import "time"

type User struct {
	Username, Email, Password	string
	CreatedAt 	time.Time
}

type Wallet struct {
	Id, UserId int
	Balance	uint
	UpdateAt, CreatedAt time.Time
}