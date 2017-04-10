package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

// UserManager struct
type UserManager struct {
	db *DB
}

// NewUserManager - Create a new *UserManager that can be used for managing users.
func NewUserManager(db *DB) (*UserManager, error) {

	db.AutoMigrate(&User{})

	usermgr := UserManager{}

	usermgr.db = db

	return &usermgr, nil
}

// HasUser - Check if the given username exists.
func (state *UserManager) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// FindUser -
func (state *UserManager) FindUser(username string) *User {
	user := User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

// FindUserByUUID -
func (state *UserManager) FindUserByUUID(uuid string) *User {
	user := User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

// AddUser - Creates a user and hashes the password
func (state *UserManager) AddUser(username, password string) *User {
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
func (state *UserManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

// CheckPassword - compare a hashed password with a possible plaintext equivalent
func (state *UserManager) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}
