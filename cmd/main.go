package main

import (
	con "gorm/connect"
	crud "gorm/crud"
	mod "gorm/models"

	"github.com/k0kubun/pp"
)

func main() {
	// connect section
	db, ConnectErr := con.ConnectToGORM("newuser", "1234", "newdb")

	if ConnectErr != nil {
		panic("Failed connect to gorm")
	}

	// automigrate section
	if aErr := con.AutoMigrate(db); aErr != nil {
		panic("Auto Migrate Error")
	}

	// Create Students Struct
	CreatedStudent := mod.Students{
		FirstName: "ali",
		LastName:  "javlonov",
		Age:       23,
		Courses: []*mod.Courses{},
	}

	// // Creare Section
	pp.Printf("\nCreate Student Func is Working\n\n")
	newStudent, csErr := crud.CreateStudent(db, &CreatedStudent)
	if csErr != nil {
		panic("Create Student Error")
	}
	pp.Printf("Created Student-> %v\n", newStudent)
	pp.Printf("Create Student Func Finished\n\n\n")

	// Get Section
	pp.Printf("\nGet Student Func is Working\n\n")
	Getstudent, gErr := crud.GetStudent(db, &newStudent)
	if gErr != nil {
		panic("Get Student Error")
	}
	pp.Println(Getstudent)
	pp.Printf("Get Student Func Finished\n\n\n")

	// Updatee Section
	pp.Printf("Update Student Func is Working\n\n")
	if uErr := crud.UpdateStudent(db, Getstudent); uErr != nil {
		panic("Updated Student Error")
	}
	pp.Println(Getstudent)
	pp.Printf("Update Student Func Finished\n\n\n")

	// Delete Section
	pp.Printf("Delete Student Func is Working\n\n")
	if dErr := crud.DeleteStudent(db, Getstudent); dErr != nil {
		panic("Delete Student Error")
	}
	pp.Println("Deleted Student")
	pp.Printf("Delete Student Func Finished\n\n\n")

	// GetAll Section
	pp.Printf("Get All Students Func is Working\n\n")
	allStudents, allErr := crud.GetAllStudents(db)
	if allErr != nil {
		panic("Erro Get All Students")
	}
	pp.Println(allStudents)
	pp.Printf("Get All Students Func Finished\n\n\n")

	pp.Println("---------------------------------------")

	// Create Courses Struct
	CreateCourse := mod.Courses{
		Name:    "QA",
		Teacher: "Skott Adkins",
		Price:   998,
		Students: []*mod.Students{
			{
				FirstName: "Murod",
				LastName:  "Otajonov",
				Age:       23,
				Courses:   []*mod.Courses{},
			},
		},
	}

	// Creare Section
	pp.Printf("\nCreate Course Func is Working\n\n")
	newCourse, ccErr := crud.CreateCourse(db, &CreateCourse)
	if ccErr != nil {
		panic("Create Course Error")
	}
	pp.Printf("Created Course--> %v\n", newCourse)
	pp.Printf("Create Course Func Finished\n\n\n")

	// Get Section
	pp.Printf("\nGet Course Func is Working\n\n")
	GetCourse, gErr := crud.GetCourse(db, &newCourse)
	if gErr != nil {
		panic("Get Course Error")
	}
	pp.Println(GetCourse)
	pp.Printf("Get Course Func Finished\n\n\n")

	// Updatee Section
	pp.Printf("Update Course Func is Working\n\n")
	if uErr := crud.UpdateCourse(db, GetCourse); uErr != nil {
		panic("Updated Course Error")
	}
	pp.Println(GetCourse)
	pp.Printf("Update Course Func Finished\n\n\n")

	// Delete Section
	pp.Printf("Delete Course Func is Working\n\n")
	if dErr := crud.DeleteCourse(db, GetCourse); dErr != nil {
		panic("Delete Course Error")
	}
	pp.Println("Deleted Course")
	pp.Printf("Delete Course Func Finished\n\n\n")

	// GetAll Section
	pp.Printf("Get All Courses Func is Working\n\n")
	allCourses, allErr := crud.GetAllCourses(db)
	if allErr != nil {
		panic("Erro Get All Courses")
	}
	pp.Println(allCourses)
	pp.Printf("Get All Courses Func Finished\n\n\n")

	pp.Println("--------------------------------------")

	// create many students
	manyStudents := []mod.Students{
		{
			FirstName: "Ali",
			LastName:  "Burxonov",
			Age:       16,
			Courses: []*mod.Courses{
				{
					Name:     "AI",
					Teacher:  "Anvar Narzulloh",
					Price:    300,
					Students: []*mod.Students{},
				},
				{
					Name:     "Java",
					Teacher:  "Alisher Kasimov",
					Price:    250,
					Students: []*mod.Students{},
				},
			},
		},
		{
			FirstName: "Jonibek",
			LastName:  "Abdullaev",
			Age:       28,
			Courses: []*mod.Courses{
				{
					Name:     "Python",
					Teacher:  "Jakhongir Rakhmonov",
					Price:    280,
					Students: []*mod.Students{},
				},
				{
					Name:     "C++",
					Teacher:  "Akbarshox Sattorov",
					Price:    270,
					Students: []*mod.Students{},
				},
			},
		},
	}
	for _, OneStudent := range manyStudents {
		oneId, csErr := crud.CreateStudent(db, &OneStudent)
		if csErr != nil {
			panic("Create Student Error")
		}
		pp.Printf("Created Student Id--> %v\n", oneId)
	}

	// create many courses
	manyCourses := []mod.Courses{
		{
			Name:    "Django",
			Teacher: "John Uik",
			Price:   345,
			Students: []*mod.Students{
				{
					FirstName: "Bekmurod",
					LastName:  "Alijonov",
					Age:       34,
					Courses:   []*mod.Courses{},
				},
				{
					FirstName: "Akmal",
					LastName:  "Kadirov",
					Age:       36,
					Courses:   []*mod.Courses{},
				},
			},
		},
		{
			Name:    "NodeJS",
			Teacher: "Saud Abdulwahed",
			Price:   450,
			Students: []*mod.Students{
				{
					FirstName: "Bahrom",
					LastName:  "Nazarov",
					Age:       37,
					Courses:   []*mod.Courses{},
				},
				{
					FirstName: "Mekmurod",
					LastName:  "Botirov",
					Age:       19,
					Courses:   []*mod.Courses{},
				},
			},
		},
	}

	for _, OneCourse := range manyCourses {
		oneId, ccErr := crud.CreateCourse(db, &OneCourse)
		if ccErr != nil {
			panic("Create Course Error")
		}
		pp.Printf("Created Course Id--> %v\n", oneId)
	}

	pp.Println("Succesfully completed all tasks without errors")
}