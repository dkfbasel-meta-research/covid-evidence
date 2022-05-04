package ictrp

import (
	"strings"
)

// randomized will check if a trial is randomized -> randomized, non-randomized, n/a
func randomized(m string) (interface{}, bool) {

	m = strings.ToLower(strings.TrimSpace(m))

	if m == "n/a" || m == "randomized" || m == "non-randomized" {
		return m, false
	}

	if strings.Contains(m, "non-randomized") {
		return "non-randomized", true
	}
	if strings.Contains(m, "non randomized") {
		return "non-randomized", true
	}
	if strings.Contains(m, "randomised: no") {
		return "non-randomized", true
	}
	if strings.Contains(m, "non-randomised") {
		return "non-randomized", true
	}
	if strings.Contains(m, "not randomized") {
		return "non-randomized", true
	}
	if strings.Contains(m, "randomized: no") {
		return "non-randomized", true
	}

	if strings.Contains(m, "randomization: not randomized") {
		return "non-randomized", true
	}

	if strings.Contains(m, "randomization: n/a") {
		return "n/a", true
	}

	if strings.Contains(m, "randomization sequence:not applic") {
		return "n/a", true
	}

	if strings.Contains(m, "randomi") {
		return "randomized", true
	}

	if strings.Contains(m, "allocation: n/a") {
		return "n/a", true
	}

	if strings.Contains(m, "masking: none") {
		return "non-randomized", true
	}

	return nil, false
}
