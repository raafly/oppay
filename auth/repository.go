package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/raafly/ewallet/helper"
	"gorm.io/gorm"
)

type AuthRepo interface{
	registerEmail(dto UserReq, token string) error
	// verifiedEmail(email, token string) error
	findByVerificationToken(verificationToken string, emailUpdate time.Time) error
	// updateToken(id, token string) error
	newUser(dto UserReq) (string, error)
	updateUser(dto UserReq) error
	findByID(id string) (*User, error)
	// getAll() (*UserReq, error)
	deleteByID(id string) error 
}

type AuthRepoImpl struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepoImpl{db: db}
}

func (r *AuthRepoImpl) registerEmail(dto UserReq, token string) error {
	err := r.db.Create(&User{ID: dto.ID, Email: dto.Email, VerificationToken: token}).Error

	if err != nil {
		log.Println(err)
		return fmt.Errorf("duplicated email")
	}

	return nil
}

func (r *AuthRepoImpl) findByVerificationToken(verificationToken string, emailUpdate time.Time) error {
	var user User
	err := r.db.Table("users").Where("verification_token = ?", verificationToken).Take(&user)
	if err != nil {
		return fmt.Errorf("token not found")
	}

	if !time.Time.IsZero(user.EmailVerifiedAt) {
		return fmt.Errorf("user already vertify")
	}

	r.db.Table("users").Where("verification_token = ?", verificationToken).Update("email_verified_at", emailUpdate)
	return nil
}

func (r *AuthRepoImpl) newUser(dto UserReq) (string, error) {
	token := helper.NewToken()
	err := r.db.Create(&User{
		Name: dto.Name,
		Pin: dto.Pin,
		RememberToken: token,
		UpdatedAt: dto.UpdatedAt,
	})

	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func (r *AuthRepoImpl) updateUser(dto UserReq) error {
	err := r.db.Where("id = ?", dto.ID).Save(&User{
		Name: dto.Name,
		Pin: dto.Pin,
		Email: dto.Email,
	}).Error

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed update")
	}

	return nil
}
func (r *AuthRepoImpl) findByID(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, fmt.Errorf("id not found")
	}

	return &user, nil
}

func (r *AuthRepoImpl) deleteByID(id string) error  {
	err := r.db.Delete(&User{}, id).Error
	if err != nil {
		return fmt.Errorf("id not found")
	}

	return nil
}

