package utils

import (
	"strconv"
)

// ParseInt64 parses a string into an int64, panicking on error.
func ParseInt64(s string) int64 {
	end, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return end
}
