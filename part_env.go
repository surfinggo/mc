package mc

import "os"

// EnvOr retrieves the value of the environment variable named
// by the key.
func EnvOr(key string, defaultValue string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return defaultValue
}
