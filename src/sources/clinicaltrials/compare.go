package clinicaltrials

import (
	"fmt"
	"strings"

	"dkfbasel.ch/covid-evidence/sources"
	"dkfbasel.ch/covid-evidence/stores"
)

// convertRecords will convert the clinicaltrials.gov records to covebasic
func convertRecords(store stores.IStore, screeningRecords []stores.Record,
	basicRecords []stores.Record, basicIndex stores.Index) ([]*stores.Record, []*stores.Record) {

	sourceName := "clinicaltrials.gov"

	// initialize the updates/inserts
	addRecords := []*stores.Record{}
	updateRecords := []*stores.Record{}

	for _, s := range screeningRecords {

		sourceID := s.Field("nct_id")
		manuallyIncluded := s.Field("cove_screening") == "include"

		if strings.Contains(strings.ToLower(s.Field("cove_screening")), "exclu") {
			continue
		}

		if !strings.Contains(strings.ToLower(s.Field("study_type")), "interven") && !manuallyIncluded {
			continue
		}

		// skip all records that exist in ninox already
		_, alreadyInBasic := basicIndex.Get(sourceID)

		// get the corresponding record if already in covebasic
		var r stores.Record
		recordFound := false
		if alreadyInBasic {
			for _, record := range basicRecords {
				if record.Field("source_id") == sourceID {
					r = record
					recordFound = true
					break
				}
			}
		}

		// if the record should be in the cove basic table but not found, skip it
		if alreadyInBasic && !recordFound && !manuallyIncluded {
			continue
		}

		if alreadyInBasic && r.Field("source") != "clinicaltrials.gov" {
			continue
		}

		// if it is a new record initialize one
		if !alreadyInBasic {

			// initialize a new record
			r = stores.Record{}
			r.Fields = make(map[string]interface{})

			r.Fields["source"] = sourceName
			r.Fields["source_id"] = sourceID

			r.Fields["review_status"] = "prefilled automatically"
			r.Fields["is_covid"] = "yes"
			r.Fields["is_trial"] = "yes"
			r.Fields["is_rct"] = "yes"
			r.Fields["is_observational"] = "no"
			r.Fields["is_duplicate"] = "no"
			r.Fields["is_rct"] = "yes"

			r.Fields["id"] = ""
		}

		r.Update("entry_type", "registration", nil)

		r.Update("url", s.Field("nct_id"), func(value string) (interface{}, bool) {
			return fmt.Sprintf("https://clinicaltrials.gov/ct2/show/record/%s", value), true
		})

		r.Update("title", s.Fields["official_title"], nil)

		abstract := ""
		if !isEmpty(s.Fields["brief_summary"]) {
			abstract = fmt.Sprintf("Brief summary:\n%s", asString(s.Fields["brief_summary"]))
		}
		if !isEmpty(s.Fields["detailed_description"]) {
			if abstract != "" {
				abstract = fmt.Sprintf("%s\n\n", abstract)
			}
			abstract = fmt.Sprintf("%s\n\nDetailed descriptions:\n%s", abstract,
				asString(s.Fields["detailed_description"]))
		}

		abstract = strings.TrimSpace(abstract)

		r.Update("abstract", abstract, nil)

		r.Update("authors", "na", nil)
		r.Update("journal", "na", nil)
		r.Update("doi", "na", nil)

		r.Update("status", s.Fields["status"], func(status string) (interface{}, bool) {
			status = strings.ToLower(status)
			if status == "not recruiting" {
				return "not yet recruiting", true
			}
			return status, false
		})

		r.Update("status_date", s.Fields["date_last_update_posted"], toIsoDate)

		r.Update("international", s.Field("location_country"), sources.InternationalFn)
		r.Update("countries", s.Field("location_country"), sources.CountryFn)
		r.Update("continents", s.Field("location_country"), sources.ContinentFn)
		r.Update("country", s.Fields["location_country"], func(country string) (interface{}, bool) {
			// country field may contain multiple countries separated by semicolon
			// -> use international if there are multiple countries
			// -> use the country name if it is the same multiple times
			if strings.Contains(country, ";") == false {
				return country, false
			}

			items := strings.Split(country, "; ")
			first := items[0]
			international := false
			for _, c := range items {
				if c != first {
					international = true
				}
			}
			if international {
				return "international", true
			}

			return first, true
		})

		// randomization
		r.Update("randomized", s.Fields["allocation"], randomized)

		// blinding
		r.Update("blinding", s.Fields["masking"], func(m string) (interface{}, bool) {
			if strings.HasPrefix(m, "None") {
				return "none", true
			}

			if strings.HasPrefix(m, "Double") || strings.HasPrefix(m, "Triple") || strings.HasPrefix(m, "Quadruple") {
				return "double blind", true
			}

			if strings.HasPrefix(m, "Single") {
				if strings.Contains(m, "Outcomes") {
					return "outcome only", true
				}
				return "single blind", true
			}

			return "", false
		})

		// longitudinal structure
		r.Update("longitudinal_structure", s.Fields["intervention_model"], toLowerCase)

		// n_arms, calculated by the number of arm types
		r.Update("n_arms", s.Fields["arm_group_arm_group_type"], func(m string) (interface{}, bool) {
			count := strings.Count(m, "; ") + 1
			return count, true
		})

		// n_enrollment
		r.Update("n_enrollment", s.Fields["enrollment"], toInt)

		// population_condition
		r.Update("population_condition", s.Fields["condition"], nil)

		// population_gender
		r.Update("population_gender", s.Fields["gender"], toLowerCase)

		// skip population_age (difficult from min-max age)
		// r.Update("population_age", "TODO", nil)

		// add intervention and control type
		_, interventionType, interventionName, interventionSubstance, controlType, control := cleanIntervention(s.Field("intervention_type"), s.Field("intervention_name"), s.Field("intervention_desc"))

		// intervention_type
		r.Update("intervention_type", interventionType, sources.SetGenerated)

		// intervention_name
		r.Update("intervention_name", interventionName, sources.SetGenerated)

		// intervention_substance
		r.Update("intervention_substance", interventionSubstance, sources.SetGenerated)

		// control_type
		r.Update("control_type", controlType, sources.SetGenerated)

		// control
		r.Update("control", control, sources.SetGenerated)

		// out_primary_measure
		r.Update("out_primary_measure", s.Fields["primary_outcome_measure"], nil)

		// out_primary_desc
		r.Update("out_primary_desc", s.Fields["primary_outcome_description"], nil)

		// out_primary_timeframe
		r.Update("out_primary_timeframe", s.Fields["primary_outcome_time_frame"], nil)

		r.Update("start_date", s.Fields["date_started"], toIsoDate)
		r.Update("end_date", s.Fields["date_completed"], toIsoDate)
		r.Update("registration_date", s.Fields["date_study_first_posted"], toIsoDate)

		r.Update("results_available", fmt.Sprintf("%s$$%s", sourceID, s.Fields["date_results_first_posted"]), resultsAvailable(store))

		// skip results_expected
		// r.Update["results_expected", "TODO", nil]

		// ipd_sharing
		r.Update("ipd_sharing", s.Fields["patient_data_sharing_ipd"], toLowerCase)

		// publication
		r.Update("publication", s.Fields["publications_pmid"], nil)

		// out_secondary_measure
		r.Update("out_secondary_measure", s.Fields["secondary_outcome_measure"], nil)

		// out_secondary_desc
		r.Update("out_secondary_desc", s.Fields["secondary_outcome_description"], nil)

		// out_secondary_timeframe
		r.Update("out_secondary_timeframe", s.Fields["secondary_outcome_time_frame"], nil)

		// inclusion exclusion criteria
		r.Update("inclusion_criteria", s.Fields["eligibility_criteria"], inclusionCriteriaFn)
		r.Update("exclusion_criteria", s.Fields["eligibility_criteria"], exclusionCriteriaFn)

		if r.Field("randomized") != "randomized" {
			r.Fields["is_rct"] = "no"
		}
		if r.Field("is_rct") == "no" && !alreadyInBasic && !manuallyIncluded {
			continue
		}

		// nothing to do, if the record was not changed
		if r.IsUpdated && alreadyInBasic {
			updateRecords = append(updateRecords, &r)
		} else if !alreadyInBasic {
			addRecords = append(addRecords, &r)
		}
	}

	return addRecords, updateRecords
}

func inclusionCriteriaFn(m string) (interface{}, bool) {
	inclusionCriteria := []string{}
	exclusionCriteria := []string{}

	allFound := false
	currentIndex := 0
	splitAtInclusion := "inclusion criteria"
	splitAtExclusion := "exclusion criteria"
	text := strings.ToLower(m)

	for !allFound {
		inclIndex := strings.Index(text, splitAtInclusion)
		exclIndex := strings.Index(text, splitAtExclusion)

		if inclIndex < 0 && exclIndex < 0 {
			return m, true
		}

		if inclIndex < 0 {
			exclusionCriteria = append(exclusionCriteria, m[currentIndex:])
			allFound = true
			continue
		}

		if exclIndex < 0 {
			inclusionCriteria = append(inclusionCriteria, m[currentIndex:])
			allFound = true
			continue
		}

		if inclIndex < exclIndex {
			inclusionCriteria = append(inclusionCriteria, m[0:exclIndex])
			text = text[exclIndex:]
			currentIndex = exclIndex
		}

		if exclIndex < inclIndex {
			exclusionCriteria = append(exclusionCriteria, m[0:inclIndex])
			text = text[inclIndex:]
			currentIndex = inclIndex
		}
	}

	return strings.Join(inclusionCriteria, "\n"), true
}

func exclusionCriteriaFn(m string) (interface{}, bool) {
	inclusionCriteria := []string{}
	exclusionCriteria := []string{}

	allFound := false
	currentIndex := 0
	splitAtInclusion := "inclusion criteria"
	splitAtExclusion := "exclusion criteria"
	text := strings.ToLower(m)

	for !allFound {
		inclIndex := strings.Index(text, splitAtInclusion)
		exclIndex := strings.Index(text, splitAtExclusion)

		if inclIndex < 0 && exclIndex < 0 {
			return m, true
		}

		if inclIndex < 0 {
			exclusionCriteria = append(exclusionCriteria, m[currentIndex:])
			allFound = true
			continue
		}

		if exclIndex < 0 {
			inclusionCriteria = append(inclusionCriteria, m[currentIndex:])
			allFound = true
			continue
		}

		if inclIndex < exclIndex {
			inclusionCriteria = append(inclusionCriteria, m[0:exclIndex])
			text = text[exclIndex:]
			currentIndex = exclIndex
		}

		if exclIndex < inclIndex {
			exclusionCriteria = append(exclusionCriteria, m[0:inclIndex])
			text = text[inclIndex:]
			currentIndex = inclIndex
		}
	}

	return strings.Join(exclusionCriteria, "\n"), true
}
