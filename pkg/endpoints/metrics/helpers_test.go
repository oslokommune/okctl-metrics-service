package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventHashConsistency(t *testing.T) {
	expectedHash := "EymlRAji5NzNJp66gDWUwMFBuu0akrY_QOQtHcmiI5k="

	for i := 0; i < 1000; i++ {
		withEvent := Event{
			Category: "somecategory",
			Action:   "someaction",
			Labels: map[string]string{
				"test":            "a",
				"this is amazing": "something relevant",
				"how many lines":  "of text can i generate",
				"i dont":          "think",
				"theres a":        "limit for that",
			},
		}

		assert.Equal(t, expectedHash, withEvent.Hash())
	}
}

func TestEventHash(t *testing.T) {
	testCases := []struct {
		name       string
		withEvent  Event
		expectHash string
	}{
		{
			name: "Should work with a single label",
			withEvent: Event{
				Category: "somecategory",
				Action:   "someaction",
				Labels: map[string]string{
					"test": "a",
				},
			},
			expectHash: "SOqDsWEz6oOJko4TOy7mS7NPPMxxLZbhr6lT5lA4b4I=",
		},
		{
			name: "Should work with multiple labels",
			withEvent: Event{
				Category: "somecategory",
				Action:   "someaction",
				Labels: map[string]string{
					"first":  "a",
					"second": "b",
					"third":  "c",
				},
			},
			expectHash: "-xp3xQ6-vw0TmMZflzmtJjmzx_futEcXxVLTEjbdxsU=",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			hash := tc.withEvent.Hash()

			assert.Equal(t, tc.expectHash, hash)
		})
	}
}
