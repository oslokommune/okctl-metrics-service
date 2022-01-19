package metrics

import (
	"testing"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

	"github.com/stretchr/testify/assert"
)

func TestIncrementingMetric(t *testing.T) {
	testCases := []struct {
		name string

		withDefinition types.Definition
		withEvent      Event

		expectErr string
	}{
		{
			name: "Should accept events with labels",

			withDefinition: types.Definition{
				Category: "testcategory1",
				Actions:  []types.Action{"testaction1"},
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

			withDefinition: types.Definition{
				Category: "testcategory2",
				Actions:  []types.Action{"testaction2"},
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
