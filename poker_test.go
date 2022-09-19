package main

import (
	"fmt"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				list: []string{},
			},
		},
		{
			name: "test1",
			args: args{
				list: []string{"红桃A"},
			},
		},
		{
			name: "test1",
			args: args{
				list: []string{"红桃A", "红桃2", "红桃3", "红桃4", "红桃5", "红桃6", "红桃7"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffle(tt.args.list)
			fmt.Println(got)
		})
	}
}
