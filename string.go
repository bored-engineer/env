package env

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

// String returns the string value of the environment variable named by the key.
func String(key string, fallback ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	panic(fmt.Errorf("%s: missing", key))
}

// StringSlice returns the []string value of the environment variable named by the key.
func StringSlice(key string, fallback ...[]string) []string {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToStringSliceE(value); err != nil {
			panic(fmt.Errorf("%s: %w", key, err))
		} else {
			return value
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	panic(fmt.Errorf("%s: missing", key))
}

// StringMapString returns the map[string]string value of the environment variable named by the key.
func StringMapString(key string, fallback ...map[string]string) map[string]string {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToStringMapStringE(value); err != nil {
			panic(fmt.Errorf("%s: %w", key, err))
		} else {
			return value
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	panic(fmt.Errorf("%s: missing", key))
}

// StringMapStringSlice returns the map[string][]string value of the environment variable named by the key.
func StringMapStringSlice(key string, fallback ...map[string][]string) map[string][]string {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToStringMapStringSliceE(value); err != nil {
			panic(fmt.Errorf("%s: %w", key, err))
		} else {
			return value
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	panic(fmt.Errorf("%s: missing", key))
}
