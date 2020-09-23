package main

import "gorm.io/gorm"

type UserService struct {
	us userService
}

type userValidator struct {
	uc userCache
}

type userCache struct {
	ug userGorm
}

type userGorm struct {
	db *gorm.DB
}

func main() {
	// When a requirements like cache imposes to our code
	// many changes should happen. Also, the code must
	// recompile
	gormdb := &gorm.DB{}
	UserService{
		uv: userValidator{
			uc: userCache{
				ug: userGorm{
					db: gormdb,
				},
			},
		},
	}
}
