package env

import (
	"os"
	"strconv"
	"time"
)

// Bool tries to parse an environment variable with the given key into a boolean value
// and falls back to the given fallback value if no environment variable was found.
func Bool(key string, fallback bool) (bool, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}

	return b, nil
}

// String tries to return an environment variable with the given key
// and falls back to the given fallback value if no environment variable was found.
func String(key string, fallback string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}

	return v, nil
}

// Int tries to parse an environment variable with the given key into an int value
// and falls back to the given fallback value if no environment variable was found.
func Int(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Float64 tries to parse an environment variable with the given key into a float64 value
// and falls back to the given fallback value if no environment variable was found.
func Float64(key string, fallback float64) (float64, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

// Duration tries to parse an environment variable with the given key into a time.Duration value
// and falls back to the given fallback value if no environment variable was found.
func Duration(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, err
	}

	return d, nil
}
