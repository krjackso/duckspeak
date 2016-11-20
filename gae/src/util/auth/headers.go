package auth

import (
	"net/http"
	"regexp"
)

var BearerRegex = regexp.MustCompile("Bearer (.*)")

func BearerFromHeader(h http.Header) (bearer string, ok bool) {
	authorization := h.Get("Authorization")
	if authorization == "" {
		return "", false
	}

	parts := BearerRegex.FindStringSubmatch(authorization)
	if len(parts) != 2 {
		return "", false
	}

	return string(parts[1]), true
}
