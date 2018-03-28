package main

import "testing"

func Test_downloadHtml(t *testing.T) {
	type args struct {
		url       string
		directory string
		retries   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadHtml(tt.args.url, tt.args.directory, tt.args.retries)
		})
	}
}
