package main

import (
	"fmt"
	"testing"
)

func Test_differenceValue(t *testing.T) {
	type args struct {
		list []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{list: []int64{5, 8, 10, 1, 3}},
			want: 9,
		},
		{
			name: "test2",
			args: args{list: []int64{5}},
			want: 0,
		},
		{
			name: "test1",
			args: args{list: []int64{1, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceValue(tt.args.list); got != tt.want {
				t.Errorf("differenceValue() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
