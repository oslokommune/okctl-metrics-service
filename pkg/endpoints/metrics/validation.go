package metrics

import validation "github.com/go-ozzo/ozzo-validation"

// Validate ensures an Event contains the required and valid data
func (receiver Event) Validate() error {
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
	return categoryValidator(c)
}

func (a Action) Validate() error {
	return actionValidator(a)
}
