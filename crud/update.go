package crud

import (
	mod "gorm/models"

	"gorm.io/gorm"
)

// this func updates first_name and age from students table
// set to first_name = 'John', age = 18
func UpdateStudent(db *gorm.DB, student *mod.Students) error {
	result := db.Model(&student).Updates(mod.Students{FirstName: "John", Age: 18})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// this func updates name and price from courses table
// set to name = 'New Upcoming Course', price = 999
func UpdateCourse(db *gorm.DB, course *mod.Courses) error {
	result := db.Model(&course).Updates(mod.Courses{Name: "New Upcoming Course", Price: 999})
	if result.Error != nil {
		return result.Error
	}
	return nil
}