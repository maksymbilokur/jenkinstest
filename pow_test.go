package main

import "testing"

func Test_pow(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "pos",
			args: 2,
			want: 1,
		},
		{
			name: "neg",
			args: -1,
			want: 1,
		},
		{
			name: "zero",
			args: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
