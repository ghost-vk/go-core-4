package main

import (
	"reflect"
	"testing"

	"github.com/go-core-4/09-ifaces/pkg/customer"
	"github.com/go-core-4/09-ifaces/pkg/employee"
)

func Test_maxAge(t *testing.T) {
	type args struct {
		peoples []Ager
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123",
			args: args{
				peoples: []Ager{customer.New(12), customer.New(12), employee.New(20)},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := maxAge(tt.args.peoples...), tt.want; got != want {
				t.Errorf("maxAge fail: got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAge2(t *testing.T) {
	type args struct {
		peoples []any
	}
	tests := []struct {
		name string
		args args
		want *employee.Employee
	}{
		{
			name: "123",
			args: args{
				peoples: []any{customer.New(12), customer.New(12), employee.New(20)},
			},
			want: employee.New(20),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := maxAge2(tt.args.peoples...), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("maxAge2 fail: got %v, want %v", got, tt.want)
			}
		})
	}
}
