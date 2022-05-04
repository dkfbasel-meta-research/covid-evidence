package ictrp

import (
	"fmt"
	"strings"
	"time"

	"dkfbasel.ch/covid-evidence/sources"
	"dkfbasel.ch/covid-evidence/stores"
)

// convertRecords will convert the clinicaltrials.gov records to covebasic
func convertRecords(store stores.IStore, screeningRecords []stores.Record, basicRecords []stores.Record,
	basicIndex stores.Index) ([]*stores.Record, []*stores.Record) {

	sourceName := "ICTRP"

	// initialize the updates/inserts
	addRecords := []*stores.Record{}
	updateRecords := []*stores.Record{}

	for _, s := range screeningRecords {

		sourceID := s.Field("trial_id")
		manuallyIncluded := s.Field("cove_screening") == "include"

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

		if alreadyInBasic && r.Field("source") != "ICTRP" {
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
			r.Fields["is_observational"] = "no"
			r.Fields["is_duplicate"] = "no"

			r.Fields["id"] = ""
		}

		r.Update("registration_date", s.Field("date_registration_3"), func(registrationDate string) (interface{}, bool) {
			asTime, err := time.Parse("20060102", registrationDate)
			if err == nil {
				return asTime.Format("2006-01-02"), true
			}
			return nil, false
		})

		r.Update("entry_type", "registration", nil)

		r.Update("url", s.Field("web_address"), nil)

		r.Update("title", s.Field("scientific_title"), nil)

		print := false
		r.Update("n_enrollment", s.Field("target_size"), getNEnrollmentFn(print))

		r.Update("corresp_author_lastname", s.Field("contact_lastname"), nil)
		r.Update("corresp_author_email", s.Field("contact_email"), func(email string) (interface{}, bool) {
			// email field could have multiple emails separated with semicolons
			if !strings.Contains(email, ";") {
				return email, false
			}

			emails := strings.Split(email, ";")
			filteredEmails := []string{}

			for _, emailPart := range emails {
				newEmail := strings.TrimSpace(emailPart)
				if newEmail == "" {
					continue
				}
				filteredEmails = append(filteredEmails, newEmail)
			}

			return strings.Join(filteredEmails, ";"), true
		})

		r.Update("status", s.Field("recruitment_status"), func(status string) (interface{}, bool) {
			status = strings.ToLower(status)
			if status == "not recruiting" {
				return "not yet recruiting", true
			}
			return status, false
		})
		r.Update("status_date", s.Field("last_refreshed_on"), toIsoDate)

		r.Update("international", s.Field("countries"), sources.InternationalFn)
		r.Update("countries", s.Field("countries"), sources.CountryFn)
		r.Update("continents", s.Field("countries"), sources.ContinentFn)
		r.Update("country", s.Field("countries"), func(country string) (interface{}, bool) {
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

		r.Update("randomized", s.Field("study_design"), randomized)

		r.Update("population_condition", s.Field("condition"), nil)

		// add intervention and control type
		l, interventionType, interventionName, interventionSubstance, controlType, control := cleanIntervention(sourceID, s.Field("intervention"))

		if l == 0 {
			r.Update("intervention_name", s.Field("intervention"), nil)
		} else {
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
		}

		r.Update("out_primary_measure", s.Field("primary_outcome"), nil)

		r.Update("start_date", s.Field("date_enrollement"), toIsoDate)

		// results_available if a results url is given
		r.Update("results_available", fmt.Sprintf("%s$$%s", sourceID, s.Field("results_yes_no")), resultsAvailable(store))

		r.Update("inclusion_criteria", s.Field("inclusion_criteria"), nil)
		r.Update("exclusion_criteria", s.Field("exclusion_criteria"), nil)

		if r.Field("randomized") != "randomized" && !alreadyInBasic && !manuallyIncluded {
			continue
		}
		if r.Field("randomized") != "randomized" && !manuallyIncluded {
			r.Fields["is_rct"] = "no"
		} else if manuallyIncluded {
			r.Fields["is_rct"] = "yes"
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
