package somepackage

import "testing"

func TestF(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{"some test", args{"some string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			F(tt.args.s)
		})
	}
}
