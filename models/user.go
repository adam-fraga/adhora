package models

import (
	d "github.com/adam-fraga/adhora/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUser returns a user by id
func (u *User) GetUser(db *d.DB) User {

	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&User{})
	var user User
	db.Client.First(&user, u.ID)
	return user
}

// GetUsers returns all users
func (u *User) GetUsers(db *d.DB) []User {

	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&User{})
	var users []User
	db.Client.Find(&users)
	return users
}

// SetUser sets a user
func (u *User) SetUser(user User, db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Save(&user)
}

// DeleteUser deletes a user
func (u *User) DeleteUser(db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Delete(&User{}, u.ID)
}
