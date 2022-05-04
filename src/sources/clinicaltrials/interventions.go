package clinicaltrials

import (
	"fmt"
	"strings"
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
// SCREENING-CLINICALTRIALS
// intervention_type: text
// intervention_name: text
// intervention_desc: text
type intervention struct {
	Type        string
	Name        string
	Description string
	Control     bool
}

func cleanIntervention(typeString, name, desc string) (int, string, string, string, string, string) {

	interventions := []intervention{}

	types := strings.Split(typeString, ";")
	names := strings.Split(name, ";")
	descriptions := strings.Split(desc, ";")

	for i := range types {
		currentIntervention := intervention{}

		// clean type
		currentIntervention.Type = strings.ToLower(strings.TrimSpace(types[i]))
		if strings.Contains(currentIntervention.Type, "biological") {
			currentIntervention.Type = "biological"
		} else if strings.Contains(currentIntervention.Type, "drug") {
			currentIntervention.Type = "drug"
		} else if strings.Contains(currentIntervention.Type, "diagnostic") {
			currentIntervention.Type = "diagnostic"
		} else if strings.Contains(currentIntervention.Type, "device") {
			currentIntervention.Type = "device"
		} else if strings.Contains(currentIntervention.Type, "procedure") {
			currentIntervention.Type = "procedure"
		} else if strings.Contains(currentIntervention.Type, "vacc") {
			currentIntervention.Type = "vaccine"
		} else {
			currentIntervention.Type = "other"
		}

		// clean name
		if i < len(names) {
			currentIntervention.Name = strings.TrimSpace(names[i])
			currentName := strings.ToLower(currentIntervention.Name)

			if strings.Contains(currentName, "vacc") {
				currentIntervention.Type = "vaccine"
			} else if strings.Contains(currentName, "placebo") ||
				strings.Contains(currentName, "sham") {
				currentIntervention.Type = "placebo/sham"
				currentIntervention.Control = true
			} else if strings.Contains(currentName, "standard") {
				currentIntervention.Type = "standard of care/no intervention"
				currentIntervention.Control = true
			}
		}

		// clean description
		if i < len(descriptions) {
			currentIntervention.Description = strings.TrimSpace(descriptions[i])
			currentDescription := strings.ToLower(currentIntervention.Description)

			if strings.Contains(currentDescription, "vacc") {
				currentIntervention.Type = "vaccine"
			} else if strings.Contains(currentDescription, "placebo") ||
				strings.Contains(currentDescription, "sham") {
				currentIntervention.Type = "placebo/sham"
				currentIntervention.Control = true
			} else if strings.Contains(currentDescription, "standard") {
				currentIntervention.Type = "standard of care/no intervention"
				currentIntervention.Control = true
			}
		}

		interventions = append(interventions, currentIntervention)
	}

	arms := len(interventions)
	interventionType := []string{}
	interventionName := []string{}
	interventionSubstance := []string{}
	control := []string{}
	controlType := []string{}

	for i := range interventions {
		if interventions[i].Control {
			controlType = append(controlType, interventions[i].Type)
			control = append(control, fmt.Sprintf("%s: %s", interventions[i].Name, interventions[i].Description))
		} else {
			interventionType = append(interventionType, interventions[i].Type)
			interventionName = append(interventionName, interventions[i].Name)
			interventionSubstance = append(interventionSubstance, interventions[i].Description)
		}
	}

	interventionTypeString := fmt.Sprintf(`["%s"]`, strings.Join(interventionType, `", "`))
	if len(interventionType) == 0 {
		interventionTypeString = ""
	}
	controlTypeString := fmt.Sprintf(`["%s"]`, strings.Join(controlType, `", "`))
	if len(controlType) == 0 {
		controlTypeString = ""
	}

	return arms, interventionTypeString, strings.Join(interventionName, ";"),
		strings.Join(interventionSubstance, ";"), controlTypeString,
		strings.Join(control, ";")
}
