package crud

import (
	m "gorm/models"

	"gorm.io/gorm"
)

// this func creates student on students table, student courses on courses table and
// connected student to course with ids
func CreateStudent(db *gorm.DB, student *m.Students) (m.Students, error) {
	result := db.Create(&student)
	if result.Error != nil {
		return m.Students{}, result.Error
	}
	return *student, result.Error
}

// this func creates course on courses table, course students on students table and
// connected courses to students with ids
func CreateCourse(db *gorm.DB, course *m.Courses) (m.Courses, error) {
	result := db.Create(&course)
	if result.Error != nil {
		return m.Courses{}, result.Error
	}
	return *course, result.Error
}