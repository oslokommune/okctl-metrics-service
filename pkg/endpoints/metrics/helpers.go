package metrics

import (
	"fmt"
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
