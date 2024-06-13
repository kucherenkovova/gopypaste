package xhttp

import (
	"net/http"
	"net/url"
)

// GetCookie returns the unescaped cookie provided in the request or empty string if not found.
func GetCookie(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	val, _ := url.QueryUnescape(cookie.Value)

	return val
}
