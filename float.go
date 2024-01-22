package env

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

// Float64 returns the float64 value of the environment variable named by the key.
func Float64(key string, fallback ...float64) float64 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToFloat64E(value); err != nil {
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

// Float32 returns the float32 value of the environment variable named by the key.
func Float32(key string, fallback ...float32) float32 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToFloat32E(value); err != nil {
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
