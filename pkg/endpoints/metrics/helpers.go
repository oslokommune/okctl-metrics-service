package metrics

import (
	"fmt"
	"sort"
	"strings"
)

func mapAsString(m map[string]string) string {
	parts := make([]string, len(m))
	index := 0

	for key, value := range m {
		parts[index] = key + value

		index++
	}

	sort.Strings(parts)

	return strings.Join(parts, "")
}

func validateAgent(legalAgents []string, agent string) error {
	for _, legalAgent := range legalAgents {
		if agent == legalAgent {
			return nil
		}
	}

	return fmt.Errorf("invalid agent")
}
