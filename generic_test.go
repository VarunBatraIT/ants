package main

import "testing"

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

func TestOneOfManyStrings_sample(t *testing.T) {
	tests := []struct {
		name string
		s    OneOfManyStrings
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.sample(); got != tt.want {
				t.Errorf("OneOfManyStrings.sample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleString(t *testing.T) {
	type args struct {
		s []string
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
			if got := SampleString(tt.args.s); got != tt.want {
				t.Errorf("SampleString() = %v, want %v", got, tt.want)
			}
		})
	}
}
