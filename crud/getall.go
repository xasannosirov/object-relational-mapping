package crud

import (
	mod "gorm/models"

	"gorm.io/gorm"
)

// this func gets all students list with all connected courses list
func GetAllStudents(db *gorm.DB) ([]*mod.Students, error) {
	var AllStudents []*mod.Students
	err := db.Model(&mod.Students{}).Preload("Courses").Find(&AllStudents).Error
	if err != nil {
		return []*mod.Students{}, err
	}
	return AllStudents, nil
}

// this func gets all courses list with all connected students list
func GetAllCourses(db *gorm.DB) ([]*mod.Courses, error) {
	var AllCourses []*mod.Courses
	err := db.Model(&mod.Courses{}).Preload("Students").Find(&AllCourses).Error
	if err != nil {
		return []*mod.Courses{}, err
	}
	return AllCourses, nil
}