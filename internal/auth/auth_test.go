package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	type test struct {
		header http.Header
		want   string
	}

	tests := []test{
		{header: http.Header{}, want: ""},
		{header: http.Header{"Authorization": {"ApiKey some_token"}}, want: "some_token"},
		{header: http.Header{"Authorization": {"ApiKey "}}, want: ""},
	}

	for i, tc := range tests {
		got, _ := GetAPIKey(tc.header)
		if got != tc.want {
			t.Fatalf("Test %v => expected: %v, got: %v", i, tc.want, got)
		}
	}
}
