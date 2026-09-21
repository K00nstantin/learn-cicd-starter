package auth

import (
	"net/http"
	"strings"
	"testing"
)

// type want struct {
// 	str string
// 	err error
// }

//var mailformedError = errors.New("malformed authorization header")

// func TestGetAPIKey(t *testing.T) {
// 	tests := map[string]struct {
// 		input http.Header
// 		want  want
// 	}{
// 		"empty": {input: http.Header{}, want: want{str: "", err: ErrNoAuthHeaderIncluded}},
// 		"wrong": {input: http.Header{"wrong": {"wrong header"}}, want: want{str: "", err: ErrNoAuthHeaderIncluded}},
// 		"valid": {input: http.Header{"Authorization": {"ApiKey key"}}, want: want{str: "key", err: nil}},
// 	}
// 	for name, ts := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got_str, got_err := GetAPIKey(ts.input)
// 			if got_err != ts.want.err || got_str != ts.want.str {
// 				t.Fatal()
// 			}
// 		})
// 	}
// }

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header_key    string
		header_value  string
		expected      string
		expectedError string
	}{
		"empty":   {expectedError: "no authorization header included"},
		"wrong":   {header_key: "wrong", header_value: "wrong value", expectedError: "no authorization header included"},
		"invalid": {header_key: "Authorization", header_value: "wrong value", expectedError: "malformed authorization header"},
		"valid":   {header_key: "Authorization", header_value: "ApiKey key", expected: "key", expectedError: "malformed authorization header"},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(ts.header_key, ts.header_value)
			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), ts.expectedError) {
					return
				}
				t.Errorf("Unexpected: %v\n", err)
				return
			}

			if got != ts.expected {
				t.Errorf("Unexpected: %s\n", got)
				return
			}
		})
	}
}
