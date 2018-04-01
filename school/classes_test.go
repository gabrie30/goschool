package school

import (
	"reflect"
	"testing"
)

func TestCreateClass(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Class
	}{
		{"Create Class", args{name: "Math"}, &Class{Name: "Math"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateClass(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateClass() = %v, want %v", got, tt.want)
			}
		})
	}
}
