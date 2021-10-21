package metrics

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	legalCharactersReRaw = "[a-z]"
	legalCharactersRe    = regexp.MustCompile(fmt.Sprintf("^%s+$", legalCharactersReRaw))
)

// Validate ensures an Event contains the required and valid data
func (receiver Event) Validate() error {
	err := validateLabels(receiver.Labels)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.Category, validation.Required),
		validation.Field(&receiver.Action, validation.Required),
	)
}

var (
	categoryValidator = generateCategoryValidator(
		CategoryCommandExecution,
	)
	actionValidator = generateActionValidator(
		commandExecutionActions,
	)
)

func (c Category) Validate() error {
	err := validation.Validate(c.String(), validation.Match(legalCharactersRe))
	if err != nil {
		return err
	}

	return categoryValidator(c)
}

func (a Action) Validate() error {
	err := validation.Validate(a.String(), validation.Match(legalCharactersRe))
	if err != nil {
		return err
	}

	return actionValidator(a)
}

func validateLabels(m map[string]string) error {
	if len(m) == 0 {
		return nil
	}

	if ok := legalCharactersRe.MatchString(mapAsString(m)); ok {
		return nil
	}

	return fmt.Errorf(
		"key or value of one or more labels includes illegal chars. Allowed characters are %s",
		legalCharactersReRaw,
	)
}
