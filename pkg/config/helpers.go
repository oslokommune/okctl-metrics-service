package config

import (
	"fmt"
	"os"
	"strconv"
)

func getString(key, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}

	return result
}

func getInt(key string, defaultValue int) (int, error) {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue, nil
	}

	resultAsInt, err := strconv.Atoi(result)
	if err != nil {
		return -1, fmt.Errorf("converting %s to int: %w", result, err)
	}

	return resultAsInt, nil
}
