package auth

import (
	"errors"
	"net/http"
	"strings"
)

const authPrefix = "ApiKey"

// Extracts an API key from the headers
// of an HTTP request
// Authorization: ApiKey (insert key here)
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 || vals[0] != authPrefix {
		return "", errors.New("authorization header wrong format")
	}

	return vals[1], nil
}
