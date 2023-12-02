package utils

import "testing"

func TestMaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"MaxInt(1, 2)", args{1, 2}, 2},
		{"MaxInt(2, 1)", args{2, 1}, 2},
		{"MaxInt(1, 1)", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
