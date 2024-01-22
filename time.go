package env

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cast"
)

// Time returns the time.Time value of the environment variable named by the key.
func Time(key string, fallback ...time.Time) time.Time {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToTimeE(value); err != nil {
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

// TimeInDefaultLocation returns the time.Time value of the environment variable named by the key.
func TimeInDefaultLocation(key string, loc *time.Location, fallback ...time.Time) time.Time {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToTimeInDefaultLocationE(value, loc); err != nil {
			panic(fmt.Errorf("%s: %w", key, err))
		} else {
			return value.In(loc)
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	panic(fmt.Errorf("%s: missing", key))
}

// Duration returns the time.Duration value of the environment variable named by the key.
func Duration(key string, fallback ...time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToDurationE(value); err != nil {
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
