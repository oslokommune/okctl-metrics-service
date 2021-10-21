package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIllegalLabelCharValidation(t *testing.T) {
	testCases := []struct {
		name       string
		withLabels map[string]string
		expectErr  string
	}{
		{
			name: "Should allow legal characters",
			withLabels: map[string]string{
				"test": "what",
			},
			expectErr: "",
		},
		{
			name:       "Should allow map containing no labels",
			withLabels: map[string]string{},
			expectErr:  "",
		},
		{
			name:       "Should allow nil map",
			withLabels: nil,
			expectErr:  "",
		},
		{
			name: "Should deny illegal characters in the label's value",
			withLabels: map[string]string{
				"test": "wha%t",
			},
			expectErr: "key or value of one or more labels includes illegal chars. Allowed characters are [a-z]",
		},
		{
			name: "Should deny illegal character in the label's key",
			withLabels: map[string]string{
				"te%st": "what",
			},
			expectErr: "key or value of one or more labels includes illegal chars. Allowed characters are [a-z]",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			event := Event{
				Category: "commandexecution",
				Action:   "scaffoldapplication",
				Labels:   tc.withLabels,
			}

			var errorMessage string

			err := event.Validate()
			if err != nil {
				errorMessage = err.Error()
			}

			assert.Equal(t, tc.expectErr, errorMessage)
		})
	}
}
