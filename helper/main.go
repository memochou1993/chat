package helper

import (
	"log"
	"net"
	"net/http"
	"strings"
)

// GetHost func
func GetHost(r *http.Request) string {
	// TEMP
	// host := r.URL.Query().Get("host")

	if host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err != nil {
		log.Println(err)
		return ""
	}

	return host
}
