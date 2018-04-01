package main

import "github.com/gabrie30/goschool/school"

func main() {

	s1 := school.CreateStudent("StudentOne", "studentone@gmail.com")
	s2 := school.CreateStudent("StudentTwo", "studentone@gmail.com")
	s3 := school.CreateStudent("StudentThree", "studentone@gmail.com")
	t1 := school.CreateTeacher("The Painter", "paint@school.edu")
	t2 := school.CreateTeacher("Professor X", "x@school.edu")
	c1 := school.CreateClass("Pottery")
	c2 := school.CreateClass("Science")

	school.AddTeacher(t1, c1)
	school.AddTeacher(t2, c2)
	school.AddStudent(s1, c1)
	school.AddStudent(s2, c1)
	school.AddStudent(s3, c1)
	school.AddStudent(s3, c2)
	school.AddStudent(s2, c2)
	school.PrintClass(c1)
	school.PrintClass(c2)
}
