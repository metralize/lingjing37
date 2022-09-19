package main

import (
	"fmt"
	"testing"
)

func Test_game(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "text1",
			args: args{
				n: 3,
				m: 2,
			},
		},
		{
			name: "text1",
			args: args{
				n: 1,
				m: 2,
			},
		},
		{
			name: "text1",
			args: args{
				n: 5,
				m: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := game(tt.args.n, tt.args.m)
			fmt.Println("---------------------  结果 -----------------------------")
			fmt.Println(got)
			fmt.Println(got1)
		})
	}
}
