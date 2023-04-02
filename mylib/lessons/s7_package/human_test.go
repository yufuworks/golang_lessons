package s7_package

import (
	"fmt"
	"testing"
)

// Example構造体名_メソッド名 でメソッドの凡例を作れる。
func ExamplePerson_SetHeight() {
	p := Person{"Mike", 20, 0}
	p.height = 180
	fmt.Println(p)
}

func TestPerson_SetHeight(t *testing.T) {
	type fields struct {
		Name   string
		Age    int
		height int
	}
	type args struct {
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				Name:   tt.fields.Name,
				Age:    tt.fields.Age,
				height: tt.fields.height,
			}
			p.SetHeight(tt.args.height)
		})
	}
}

func TestSay(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Say()
		})
	}
}
