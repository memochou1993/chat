package helper

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Missing .env file.")
	}
}

// GetEnv func
func GetEnv(name string) string {
	env := os.Getenv(name)

	if env == "" {
		log.Fatalf("Environment variable not found: \"%s\".", name)
	}

	return env
}

// GetPlatform func
func GetPlatform(r *http.Request) string {
	if !IsProduction() {
		return r.URL.Query().Get("platform")
	}

	return r.Header.Get("User-Agent")
}

// GetHost func
func GetHost(r *http.Request) string {
	if !IsProduction() {
		return r.URL.Query().Get("host")
	}

	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))

	if err != nil {
		log.Println(err)
		return ""
	}

	return host
}

// IsLocal func
func IsLocal() bool {
	return GetEnv("APP_ENV") == "local"
}

// IsProduction func
func IsProduction() bool {
	return GetEnv("APP_ENV") == "production"
}
