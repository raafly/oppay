package core

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo interface {
	signUp(data User) error
	singIn(data User) (*User, error)
	findById(UserId string) (*User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(DB *sql.DB) UserRepo {
	return &userRepo{db: DB}
}

func (r userRepo) signUp(data User) error {
	sql := `INSERT INTO users(username, email, password) VALUES($1, $2, $3)`
	if _, err := r.db.Exec(sql, data.Username, data.Email, data.Password); err != nil {
		return errors.New("failed create account")
	}

	sql = `INSERT INTO wallet(id, user_id, balance) VALUES(?, ?, ?)`
	
	return nil
}

func (r userRepo) singIn(data User) (*User, error) {
	sql := `SELECT username, email, password FROM users WHERE username = $1`
	rows, err := r.db.Query(sql, data.Username)
	if err != nil {
		return nil, errors.New("account not found")
	}
	defer rows.Close()

	data = User{}
	for rows.Next() {
		err := rows.Scan(&data.Username, &data.Email, &data.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to exec query %v", err)
		}
	}
	return &data, nil
}

func (r userRepo) findById(UserId string) (*User, error) {
	sql := `SELECT username, email, password, created_at FROM users WHERE username = $1`
	rows, err := r.db.Query(sql, UserId)

	if err != nil {
		return nil, errors.New("account not found")
	}
	defer rows.Close()

	data := User{}
	for rows.Next() {
		err := rows.Scan(&data.Username, &data.Email, &data.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to exec query %v", err)
		}
	}
	return &data, nil
}