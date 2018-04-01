package school

// Student is a student
type Student struct {
	Name  string
	Email string
}

// CreateStudent creates a new student
func CreateStudent(name string, email string) *Student {
	s := &Student{
		Name:  name,
		Email: email,
	}

	return s
}
