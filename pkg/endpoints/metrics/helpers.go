package metrics

import (
	"fmt"
	"sort"
	"strings"
)

func generateActionValidator(actionList ...[]Action) func(Action) error {
	combinedList := make([]Action, 0)

	for _, list := range actionList {
		combinedList = append(combinedList, list...)
	}

	return func(value Action) error {
		for _, legal := range combinedList {
			if legal == value {
				return nil
			}
		}

		return fmt.Errorf("illegal action: %s", string(value))
	}
}

func generateCategoryValidator(categories ...Category) func(Category) error {
	return func(value Category) error {
		for _, legal := range categories {
			if legal == value {
				return nil
			}
		}

		return fmt.Errorf("illegal category: %s", string(value))
	}
}

func mapAsString(m map[string]string) string {
	parts := make([]string, len(m))
	index := 0

	for key, value := range m {
		parts[index] = key + value

		index++
	}

	sort.Strings(parts)

	return strings.Join(parts, "")
}
