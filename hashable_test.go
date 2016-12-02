package main

import "testing"

func TestRemoveHashes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

		{"Remove Hashes", args{str: "this #is a te#st"}, "this is a te#st"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveHashes(tt.args.str); got != tt.want {
				t.Errorf("RemoveHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddHashes(t *testing.T) {
	type args struct {
		str   string
		total int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Remove Hashes", args{str: "this #is a leadership"}, "this is a #leadership"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddHashes(tt.args.str, tt.args.total); got != tt.want {
				t.Errorf("AddHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
