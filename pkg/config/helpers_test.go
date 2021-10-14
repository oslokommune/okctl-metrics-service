package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	testCases := []struct {
		name string

		withValue string
		expect    []string
	}{
		{
			name: "Should work with a single value",

			withValue: "okctl",
			expect:    []string{"okctl"},
		},
		{
			name: "Should work with multiple values",

			withValue: "okctl;okctl-dev;nese",
			expect:    []string{"okctl", "okctl-dev", "nese"},
		},
		{
			name: "Should work with multiple values, finishing with a semi colon",

			withValue: "okctl;okctl-dev;nese;",
			expect:    []string{"okctl", "okctl-dev", "nese"},
		},
		{
			name: "Should work with no values",

			withValue: "",
			expect:    []string{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			getter := func(_ string) string {
				return tc.withValue
			}

			result := getStringList(getter, "", []string{})

			assert.Equal(t, tc.expect, result)
		})
	}
}
