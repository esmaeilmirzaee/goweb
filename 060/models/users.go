package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("models: resources not found.")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
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

// Close method closes user service database connection.
func (us *UserService) Close() error {
	return nil
	// us.db.Close()
}

func (us *UserService) DestructiveReset(hardRest bool) error {
	if hardRest {
		return us.db.Migrator().DropTable(&User{})
	} else {
		return us.db.AutoMigrate(&User{})
	}
}
