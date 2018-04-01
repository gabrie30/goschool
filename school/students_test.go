package school

import (
	"reflect"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want *Student
	}{
		{"First Test", args{name: "Foobar", email: "email.com"}, &Student{Name: "Foobar", Email: "email.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateStudent(tt.args.name, tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
