package metrics

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var legalCharactersRe = regexp.MustCompile(`[a-z]+`)

// Validate ensures an Event contains the required and valid data
func (receiver Event) Validate() error {
	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.Category, validation.Required),
		validation.Field(&receiver.Action, validation.Required),
		validation.Field(&receiver.Label, validation.Match(legalCharactersRe)),
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
