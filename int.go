package env

import (
	"fmt"
	"os"

	"github.com/spf13/cast"
)

// Int returns the integer value of the environment variable named by the key.
func Int(key string, fallback ...int) int {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToIntE(value); err != nil {
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

// Int16 returns the int16 value of the environment variable named by the key.
func Int16(key string, fallback ...int16) int16 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToInt16E(value); err != nil {
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

// Int32 returns the int32 value of the environment variable named by the key.
func Int32(key string, fallback ...int32) int32 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToInt32E(value); err != nil {
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

// Int64 returns the int64 value of the environment variable named by the key.
func Int64(key string, fallback ...int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToInt64E(value); err != nil {
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

// Int8 returns the int8 value of the environment variable named by the key.
func Int8(key string, fallback ...int8) int8 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToInt8E(value); err != nil {
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

// Uint returns the uint value of the environment variable named by the key.
func Uint(key string, fallback ...uint) uint {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToUintE(value); err != nil {
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

// Uint16 returns the uint16 value of the environment variable named by the key.
func Uint16(key string, fallback ...uint16) uint16 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToUint16E(value); err != nil {
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

// Uint32 returns the uint32 value of the environment variable named by the key.
func Uint32(key string, fallback ...uint32) uint32 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToUint32E(value); err != nil {
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

// Uint64 returns the uint64 value of the environment variable named by the key.
func Uint64(key string, fallback ...uint64) uint64 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToUint64E(value); err != nil {
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

// Uint8 returns the uint8 value of the environment variable named by the key.
func Uint8(key string, fallback ...uint8) uint8 {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := cast.ToUint8E(value); err != nil {
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
