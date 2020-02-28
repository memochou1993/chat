package helper

import (
	"net"
	"net/http"
	"strings"
)

// GetHost func
func GetHost(r *http.Request) string {
	// host := r.URL.Query().Get("host")

	// return host

	if host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return host
	}

	return ""
}
