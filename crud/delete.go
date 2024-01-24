package crud

import (
	m "gorm/models"

	"gorm.io/gorm"
)

// this func disconnect courses with student from student_course table and
// delete student from students table
func DeleteStudent(db *gorm.DB, student *m.Students) error {
	result := db.Exec("DELETE FROM student_course WHERE students_id = $1", student.Id)
	if result.Error != nil {
		return result.Error
	}
	Dresult := db.Delete(student)
	if Dresult.Error != nil {
		return Dresult.Error
	}
	return nil
}

// this func disconnect students with course from student_course table and
// delete course from courses table
func DeleteCourse(db *gorm.DB, course *m.Courses) error {
	result := db.Exec("DELETE FROM student_course WHERE courses_id = $1", course.Id)
	if result.Error != nil {
		return result.Error
	}
	Dresutl := db.Delete(course)
	if Dresutl.Error != nil {
		return Dresutl.Error
	}
	return nil
}