package main

import (
	"reflect"
	"testing"
)

func Test_expandUrl(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []error
	}{
		// TODO: Add test cases.
		{"feed proxy", args{"http://feedproxy.google.com/~r/Unbounce/~3/E9tAubzJeMk/"}, "http://unbounce.com/landing-pages/how-big-brands-use-urgency-to-drive-conversions-during-the-holidays/?utm_source=Social&utm_medium=RSS", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := expandUrl(tt.args.uri)
			if got != tt.want {
				t.Errorf("expandUrl() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("expandUrl() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_in_array(t *testing.T) {
	type args struct {
		v  interface{}
		in interface{}
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
		wantI  int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotI := in_array(tt.args.v, tt.args.in)
			if gotOk != tt.wantOk {
				t.Errorf("in_array() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotI != tt.wantI {
				t.Errorf("in_array() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func Test_alphanumericsmall(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphanumericsmall(tt.args.str); got != tt.want {
				t.Errorf("alphanumericsmall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collectKeys(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collectKeys(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleString(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SampleString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SampleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}
