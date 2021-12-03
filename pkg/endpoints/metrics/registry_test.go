package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementingMetric(t *testing.T) {
	testCases := []struct {
		name string

		withDefinition Definition
		withEvent      Event

		expectErr string
	}{
		{
			name: "Should accept events with labels",

			withDefinition: Definition{
				Category: "testcategory1",
				Actions:  []Action{"testaction1"},
				Labels:   []string{"a"},
			},
			withEvent: Event{
				Category: "testcategory1",
				Action:   "testaction1",
				Labels: map[string]string{
					"a": "b",
				},
			},
		},
		{
			name: "Should accept events without labels",

			withDefinition: Definition{
				Category: "testcategory2",
				Actions:  []Action{"testaction2"},
				Labels:   nil,
			},
			withEvent: Event{
				Category: "testcategory2",
				Action:   "testaction2",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			registry := NewMetricRegistry()

			registry.Add(tc.withDefinition)

			err := registry.Increment("okctldev", tc.withEvent)

			if len(tc.expectErr) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.expectErr, err.Error())
			}

			registry.Reset()
		})
	}
}
