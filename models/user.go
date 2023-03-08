package models

import (
	
	"gorm.io/gorm"
"time"	
)

// User struct
type User struct {
	ID uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:100;not null" json:"-"`
	Password  string    `gorm:"size:100;not null;" json:"-"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"-"`
    UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"-"`
}
func CreateUser(db *gorm.DB,User *User)  (error,bool) {
	result := db.Create(User)
	err := result.Error
	if err != nil {
		return err,false
	}
	return err,true
 }
 
 //get users
 func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
	   return err
	}
	return nil
 }
 
 //get user by id
 func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
	   return err
	}
	return nil
 }
 
 //update user
 func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
 }
 
 //delete user
 func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
 }