package school

// Teacher is a teacher
type Teacher struct {
	Name  string
	Email string
}

// CreateTeacher creates a new eacher
func CreateTeacher(name string, email string) *Teacher {
	t := &Teacher{
		Name:  name,
		Email: email,
	}

	return t
}
