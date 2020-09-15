package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("models: resources not found.")
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

type UserService struct {
	db *gorm.DB
}

// NewUserService returns a database connection
//
//
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &UserService{
		db: db,
	}, nil
}

// ByID returns user
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create fills table
func (us *UserService) Create(user *User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	// return nil
	return us.db.Create(user).Error
}

// Close method closes user service database connection.
func (us *UserService) Close() error {
	return nil
	// us.db.Close()
}

func (us *UserService) DestructiveReset() error {
	if err := us.db.Migrator().DropTable(&User{}).Error; err != nil {
		return err
	}
	us.AutoMigrate(&User{})
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&user{}); err != nil {
		return err
	}
	return nil
}
