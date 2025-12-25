package env

import (
	"log/slog"
	"os"
)

func GetString(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	slog.Error("missing env variable", "variable", key)
	os.Exit(1)

	return ""
}
