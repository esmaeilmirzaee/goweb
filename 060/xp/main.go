package main

import "gorm.io/gorm"

type User struct {
	id   uint64
	Name string
}

type UserReader interface {
	ByID(id uint) (*User, error)
}

type UserService struct {
	UserReader
}

type userValidator struct {
	UserReader
}

type userCache struct {
	UserReader
}

// userGorm is the bottom layer so it should be implement
// the UserReader...
type userGorm struct {
	db *gorm.DB
}

func (ug userGorm) ByID(id uint64) (*User, error) {
	var user User
	user = User{id: 1233127462314673264, Name: "a a"} // db.Where...
	return &user, nil
}

func main() {
	gormdb := &gorm.DB{}
	us := UserService{
		UserReader: userValidator{
			UserReader: userGorm{
				db: gormdb,
			},
		},
	}
}
