package crud

import (
	mod "gorm/models"

	"gorm.io/gorm"
)

// this func gets student with all connected courses
func GetStudent(db *gorm.DB, student *mod.Students) (*mod.Students, error) {
	var NewStudent mod.Students
	res := db.Model(mod.Students{Id: student.Id}).First(&NewStudent)
	if res.Error != nil {
		return &mod.Students{}, res.Error
	}
	return &NewStudent, nil
}

// this func gets course with all connected students
func GetCourse(db *gorm.DB, course *mod.Courses) (*mod.Courses, error) {
	var NewCourse mod.Courses
	res := db.Model(mod.Courses{Id: course.Id}).First(&NewCourse)
	if res.Error != nil {
		return &mod.Courses{}, res.Error
	}
	return &NewCourse, nil
}