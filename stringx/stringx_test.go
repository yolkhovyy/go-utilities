package stringx

import (
	"testing"
)

func TestCapitalize(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				in: "",
			},
			want: "",
		},
		{
			name: "lower",
			args: args{
				in: "lower",
			},
			want: "Lower",
		},
		{
			name: "Upper",
			args: args{
				in: "Upper",
			},
			want: "Upper",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.in); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSpaceToLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "space",
			args: args{
				s: " ",
			},
			want: "",
		},
		{
			name: "upper",
			args: args{
				s: "UPPER",
			},
			want: "upper",
		},
		{
			name: "lower",
			args: args{
				s: "lower",
			},
			want: "lower",
		},
		{
			name: "mixed",
			args: args{
				s: " MiXeD ",
			},
			want: "mixed",
		},
		{
			name: "whitespace",
			args: args{
				s: "  \n\r\t\f\vwhitespace\n\r\t\f\v  ",
			},
			want: "whitespace",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpaceToLower(tt.args.s); got != tt.want {
				t.Errorf("TrimSpaceToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
