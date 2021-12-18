package dao

import "testing"

func TestInsertOneArticle(t *testing.T) {
	type args struct {
		subject string
		url     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertOneArticle(tt.args.subject, tt.args.url)
		})
	}
}
