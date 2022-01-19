package types

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	legalCharactersReRaw = "[a-z]"
	legalCharactersRe    = regexp.MustCompile(fmt.Sprintf("^%s+$", legalCharactersReRaw))
)

func (c Category) String() string {
	return string(c)
}

func (c Category) Validate() error {
	return validation.Validate(c.String(), validation.Match(legalCharactersRe))
}

func (a Action) String() string {
	return string(a)
}

func (a Action) Validate() error {
	return validation.Validate(a.String(), validation.Match(legalCharactersRe))
}
