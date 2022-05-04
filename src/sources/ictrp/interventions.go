package ictrp

import (
	"strings"

	"dkfbasel.ch/covid-evidence/sources/ictrp/interventions"
)

// cleanIntervention will split the intervention field into
// classification types
//
// COVE_BASIC
// intervention_type: ["drug", "vaccine", "biological", "traditional medicine", "device", "diagnostic", "procedure", "other", "unclear"]
// intervention_name: text
// intervention_substance: text
// control: text
// control_type: ["active control", "placebo/sham", "standard of care/no intervention", "no control", "waiting list", "other control", "unclear"]
//
// SCREENING-ICTRP
// intervention: text
func cleanIntervention(sourceID, intervention string) (int, string, string, string, string, string) {

	if strings.Contains(sourceID, "EUCTR") {
		return interventions.Euro(intervention)
	}

	return 0, "", "", "", "", ""
}
