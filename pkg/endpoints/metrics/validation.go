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
