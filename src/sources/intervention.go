package sources

import (
	"regexp"
	"strings"

	"dkfbasel.ch/covid-evidence/stores"
	"dkfbasel.ch/covid-evidence/logger"
)

// Interventions get intervention type and name out of the intervention field
func Interventions(store stores.IStore, coveID int32, interventionValue string) error {

	isStructuredFn := func(intervention string) bool {
		if strings.HasPrefix(intervention, "Drug: ") ||
			strings.HasPrefix(intervention, "Combination Product: ") ||
			strings.HasPrefix(intervention, "Other: ") ||
			strings.HasPrefix(intervention, "Biological: ") ||
			strings.HasPrefix(intervention, "Dietary Supplement: ") ||
			strings.HasPrefix(intervention, "Device: ") ||
			strings.HasPrefix(intervention, "Procedure: ") {
			return true
		}
		return false
	}

	// check if intervetion field is structured
	isStructured := isStructuredFn(interventionValue)

	if !isStructured {
		return store.FindExistingInterventions(coveID, interventionValue)
	}

	// list of interventions splitted
	interventions := []string{}

	if strings.Contains(interventionValue, ";") && !strings.Contains(interventionValue, "\n") {
		interventions = strings.Split(interventionValue, ";")
	} else if strings.Contains(interventionValue, "\n") {
		interventions = strings.Split(interventionValue, "\n")
	} else {
		return nil
	}

	interventionMap := make([]map[string]string, 0)

	// go thru all interventions skip a line if it does not start with an intervention
	// type
	for _, intervention := range interventions {
		// the line does not starts with a intervention type: skip it
		if !isStructuredFn(intervention) {
			continue
		}

		// split the prefix from the intervention name
		interventionParts := strings.Split(intervention, ": ")

		// if only the prefix is in the current line skip it
		if len(interventionParts) != 2 {
			continue
		}

		// The left side is the intervention type and the right side
		// is the intervention name
		interventionType := strings.ToLower(strings.TrimSpace(interventionParts[0]))
		interventionName := strings.TrimSpace(interventionParts[1])

		// remove quantity if available
		quantityRegex := regexp.MustCompile(`\d+\s?(mg|MG)`)
		quantityIndex := quantityRegex.FindStringIndex(interventionName)
		if len(quantityIndex) == 2 && quantityIndex[0] > 0 {
			interventionName = interventionName[0:(quantityIndex[0] - 1)]
		}

		// adapt the intervention type if some keywords are in the intervention name
		lowerInterventionName := strings.TrimSpace(strings.ToLower(interventionName))
		if strings.Contains(lowerInterventionName, "vacc") {
			interventionType = "vaccine"
		}
		if strings.HasPrefix(lowerInterventionName, "standard") &&
			(strings.Contains(lowerInterventionName, "treatment") ||
				strings.Contains(lowerInterventionName, "care") ||
				strings.Contains(lowerInterventionName, "therapy") ||
				strings.Contains(lowerInterventionName, "procedure") ||
				strings.Contains(lowerInterventionName, "strategy") ||
				strings.Contains(lowerInterventionName, "system")) {
			interventionType = "standard of care"
		}
		if interventionType != "standard of care" && strings.HasPrefix(lowerInterventionName, "placebo") {
			interventionType = "placebo"
		} else if interventionType == "standard of care" && strings.Contains(lowerInterventionName, "placebo") {
			interventionType = "placebo"
		}

		// if it is not standard of care but nevertheless standard of care is mentioned remove it
		if interventionType != "standard of care" {
			standardCare := regexp.MustCompile(`(\+\s?)?(standard\sof\scare)|(standard\streatment)|(standard\scare)`)
			standardIndex := standardCare.FindStringIndex(lowerInterventionName)
			if len(standardIndex) == 2 && standardIndex[0] > 0 {
				interventionName = interventionName[0:(standardIndex[0] - 1)]
			}
		}

		newIntervention := make(map[string]string)
		newIntervention["type"] = interventionType
		newIntervention["name"] = interventionName
		interventionMap = append(interventionMap, newIntervention)
	}

	err := store.SaveIntervention(coveID, interventionMap)
	if err != nil {
		logger.Error("could save intervention", err)
		return err
	}

	return nil
}
