package config

import (
	"fmt"
	"strconv"
	"strings"
)

func getString(getter stringValueGetter, key, defaultValue string) string {
	result := getter(key)
	if result == "" {
		return defaultValue
	}

	return result
}

func getInt(getter stringValueGetter, key string, defaultValue int) (int, error) {
	result := getter(key)
	if result == "" {
		return defaultValue, nil
	}

	resultAsInt, err := strconv.Atoi(result)
	if err != nil {
		return -1, fmt.Errorf("converting %s to int: %w", result, err)
	}

	return resultAsInt, nil
}

func getStringList(getter stringValueGetter, key string, defaultValue []string) []string {
	result := getter(key)
	if result == "" {
		return defaultValue
	}

	if strings.HasSuffix(result, ";") {
		result = result[:len(result)-1]
	}

	return strings.Split(result, ";")
}
