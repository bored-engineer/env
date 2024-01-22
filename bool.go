package env

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

// Bool returns the boolean value of the environment variable named by the key.
func Bool(key string, fallback ...bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToBoolE(value); err != nil {
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
