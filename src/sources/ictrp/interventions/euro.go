package interventions

import (
	"fmt"
	"strings"
)

func Euro(interventionSource string) (int, string, string, string, string, string) {
	interventions := strings.Split(interventionSource, "<br><br>")

	interventionTypes := []string{}
	interventionSubstances := []string{}
	interventionNames := []string{}
	controlTypes := []string{}
	controls := []string{}

	for i := range interventions {
		intervention := interventions[i]

		lines := strings.Split(intervention, "<br>")

		interventionType := "unclear"
		if strings.Contains(strings.ToLower(intervention), "vacc") {
			interventionType = "vaccine"
		}

		substance := ""
		name := []string{}

		controlType := ""
		control := []string{}

		for l := range lines {
			line := strings.TrimSpace(lines[l])

			if line == "" {
				continue
			}

			if line == "," || line == ":" || line == ";" || len(line) < 8 {
				continue
			}

			if substance == "" {
				keyValue := strings.Split(line, ":")
				if len(keyValue) != 2 {
					continue
				}

				substance = strings.TrimSpace(keyValue[1])
				continue
			}
			if strings.Contains(strings.ToLower(line), "placebo") {
				if controlType == "" {
					controlType = "placebo/sham"
				}
				control = append(control, strings.TrimSpace(line))
				continue
			}

			name = append(name, line)
		}

		interventionTypes = append(interventionTypes, interventionType)
		interventionSubstances = append(interventionSubstances, substance)
		interventionNames = append(interventionNames, strings.Join(name, ","))
		controlTypes = append(controlTypes, controlType)
		controls = append(controls, strings.Join(control, ","))
	}

	interventionTypeString := fmt.Sprintf(`["%s"]`, strings.Join(interventionTypes, ","))
	if len(interventionTypes) == 0 {
		interventionTypeString = ""
	}

	controlTypeString := fmt.Sprintf(`["%s"]`, strings.Join(controlTypes, ","))
	if len(controlTypes) == 0 {
		interventionTypeString = ""
	}

	return len(interventionTypes), interventionTypeString, strings.Join(interventionNames, ";"),
		strings.Join(interventionSubstances, ";"), controlTypeString, strings.Join(controls, ";")
}
