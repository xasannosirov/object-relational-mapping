package models

// struct to work with students table
type Students struct {
	Id        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Age       int
	Courses   []*Courses `gorm:"many2many:student_course;"`
}

// struct to work with courses table
type Courses struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Teacher  string
	Price    int
	Students []*Students `gorm:"many2many:student_course;"`
}