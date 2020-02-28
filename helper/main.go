package helper

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload" // initialize
)

var (
	env = os.Getenv("APP_ENV")
)

// GetHost func
func GetHost(r *http.Request) string {
	if env == "local" {
		return r.URL.Query().Get("host")
	}

	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

	if err != nil {
		log.Println(err)
		return ""
	}

	return host
}
