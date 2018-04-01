package school

import "fmt"

//
type Class struct {
	Name     string
	Teacher  Teacher
	Students []Student
}

//
func CreateClass(name string) *Class {
	c := &Class{
		Name: name,
	}

	return c
}

// AddTeacher adds a teacher to a class
func AddTeacher(t *Teacher, c *Class) {
	c.Teacher = *t
}

//
func AddStudent(s *Student, c *Class) {
	c.Students = append(c.Students, *s)
}

// ListStudents pretty prints all students in a class
func ListStudents(c *Class) {
	fmt.Println("The students in class, ", c.Name)
	for _, s := range c.Students {
		fmt.Println(s.Name)
	}
}

// PrintClass prints the entire class
func PrintClass(c *Class) {
	fmt.Println()
	fmt.Println("------ Class " + c.Name + " -------")
	fmt.Println()
	fmt.Println("Teacher: ", c.Teacher.Name)
	fmt.Println()
	fmt.Println("------ Students ------------")
	fmt.Println()
	for _, s := range c.Students {
		fmt.Println(s.Name)
	}
	fmt.Println()
}
