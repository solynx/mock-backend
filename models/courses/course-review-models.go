package courses

import (
	"gorm.io/gorm"

)
type Review struct {
	Id uint32    `gorm:"primary_key;auto_increment" json:"-"`
	Star float32	 `json:"star"`
	CourseId uint32 `json:"-"`
	UserId uint32 `json:"-"`
	Total uint32	 `json:"total"`
}
// New Review
func CreateReview (db *gorm.DB,Review *Review) (err error){
	err =db.Create(Review).Error
	if err != nil {
		return err
	}
	return nil
}
//Get all Reviews
func GetReviews(db *gorm.DB,Review []*Review) (err error) {
	err = db.Find(Review).Error
	if err != nil {
		return err
	 }
	 return nil
}
//get a Review
func GetReviewById(db *gorm.DB, Review *Review, id string) (err error) {
	err = db.Where("id = ?", id).First(Review).Error
	if err != nil {
	   return err
	}
	return nil
 }
//update a Review
 func UpdateReview(db *gorm.DB, Review *Review) (err error) {
	db.Save(Review)
	return nil
 }
 // delete a Review
 func DeleteReview(db *gorm.DB, Review *Review, id int) (err error) {
	db.Where("id = ?", id).Delete(Review)
	return nil
 }