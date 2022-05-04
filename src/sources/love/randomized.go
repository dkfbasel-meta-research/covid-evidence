package love

import (
	"strings"
)

// randomized will check if a trial is randomized -> randomized, non-randomized, n/a
func randomized(m string) (interface{}, bool) {

	m = strings.ToLower(strings.TrimSpace(m))

	if m == "n/a" || m == "randomized" || m == "non-randomized" {
		return m, false
	}

	if strings.Contains(m, "randomized: no") {
		return "non-randomized", true
	}

	if strings.Contains(m, "randomised: yes") {
		return "randomized", true
	}

	if strings.Contains(m, "randomization: randomized") {
		return "randomized", true
	}

	if strings.Contains(m, "non-randomized clinical trial") {
		return "non-randomized", true
	}

	if strings.Contains(m, "randomized controlled trial") {
		return "randomized", true
	}

	if strings.Contains(m, "non-randomized-controlled") {
		return "non-randomized", true
	}

	if strings.Contains(m, "randomized-controlled") {
		return "randomized", true
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
