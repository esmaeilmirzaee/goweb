package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound        = errors.New("models: resources not found.")
	ErrInvalidId       = errors.New("models: invalid id is provided.")
	ErrInvalidEmail    = errors.New("models: Email is invalid.")
	ErrInvalidPassword = errors.New("models: invalid password provided.")
)

const userPwPepper = "bookish-umbrella-blissful"

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

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	switch err {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return ErrNotFound
	default:
		return err
	}
}

// ByID returns user
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail checks the existance of an account in the database.
//
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Create fills table
func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	// return nil
	return us.db.Create(user).Error
}

// Authenticate can be authenticated a user by provided email and
// password.
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	pwByte := []byte(password + userPwPepper)
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), pwByte)
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}
	cookie := http.Cookie(
		Name "email",
		Value: foundUser.Email
	)
	http.SetCookie(&cookie)


	return foundUser, nil
}

// Close method closes user service database connection.
func (us *UserService) Close() error {
	return nil
	// us.db.Close()
}

// DestructiveReset drops the table and recreates it.
func (us *UserService) DestructiveReset() {
	us.db.Migrator().DropTable(&User{})
	us.AutoMigrate()
}

// AutoMigrate restalblish user tables in the database.
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

// Update user
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete user with the provided id.
//
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}
