package models

import (
	"github.com/jinzhu/gorm"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	UUID     string `gorm:"not null;unique"`
}

// UserState struct
type UserState struct {
	db *gorm.DB
}

// NewUserState - Create a new *UserState that can be used for managing users.
func NewUserState(db *gorm.DB) (*UserState, error) {

	// Test connection
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	state := new(UserState)

	state.db = db

	return state, nil
}

// HasUser - Check if the given username exists.
func (state *UserState) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// FindUser -
func (state *UserState) FindUser(username string) *User {
	user := User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

// FindUserByUUID -
func (state *UserState) FindUserByUUID(uuid string) *User {
	user := User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

// AddUser - Creates a user and hashes the password
func (state *UserState) AddUser(username, password string) *User {
	passwordHash := state.HashPassword(username, password)
	user := &User{
		Username: username,
		Password: passwordHash,
		UUID:     uuid.NewV4().String(),
	}
	state.db.Create(&user)
	return user
}

// HashPassword - Hash the password (takes a username as well, it can be used for salting).
func (state *UserState) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}
