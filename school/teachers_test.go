package school

import (
	"reflect"
	"testing"
)

func TestCreateTeacher(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want *Teacher
	}{
		{"Create Teacher", args{name: "foo", email: "bar.com"}, &Teacher{"foo", "bar.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTeacher(tt.args.name, tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTeacher() = %v, want %v", got, tt.want)
			}
		})
	}
}
