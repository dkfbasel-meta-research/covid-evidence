package clinicaltrials

import (
	"strings"
)

// randomized will check if a trial is randomized -> randomized, non-randomized, n/a
func randomized(m string) (interface{}, bool) {

	m = strings.ToLower(strings.TrimSpace(m))

	if m == "n/a" || m == "randomized" || m == "non-randomized" {
		return m, false
	}

	if strings.Contains(m, "allocation: randomized") {
		return "randomized", true
	}

	if strings.Contains(m, "allocation: n/a") {
		return "n/a", true
	}

	if strings.Contains(m, "allocation: non-randomized") {
		return "non-randomized", true
	}

	if strings.Contains(m, "masking: none") {
		return "non-randomized", true
	}

	return nil, false
}
