package utils

import (
	"fmt"
	"strconv"
)

func StringToUint(s string) (uint, error) {
	// Convert string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("conversion error: %w", err)
	}

	// Check if the integer is non-negative and within uint range
	if i < 0 || i > int(^uint(0)>>1) {
		return 0, fmt.Errorf("value out of range for uint")
	}

	return uint(i), nil
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// Handle error, return default value, or propagate error as needed
		return 0 // Return 0 as default value in case of error
	}
	return i
}
