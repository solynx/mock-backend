package courses

import (
	"gorm.io/gorm"
	"time"
)
type Enrollment struct {
	Id uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserId uint32
	CourseId uint32
	Status bool
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
// New Enrollment
func CreateEnrollment (db *gorm.DB,Enrollment *Enrollment) (err error){
	err =db.Create(Enrollment).Error
	if err != nil {
		return err
	}
	return nil
}
//Get all Enrollments
func GetEnrollments(db *gorm.DB,Enrollment []*Enrollment) (err error) {
	err = db.Find(Enrollment).Error
	if err != nil {
		return err
	 }
	 return nil
}
//get a Enrollment
func GetEnrollmentById(db *gorm.DB, Enrollment *Enrollment, id string) (err error) {
	err = db.Where("id = ?", id).First(Enrollment).Error
	if err != nil {
	   return err
	}
	return nil
 }
//update a Enrollment
 func UpdateEnrollment(db *gorm.DB, Enrollment *Enrollment) (err error) {
	db.Save(Enrollment)
	return nil
 }
 // delete a Enrollment
 func DeleteEnrollment(db *gorm.DB, Enrollment *Enrollment, id int) (err error) {
	db.Where("id = ?", id).Delete(Enrollment)
	return nil
 }