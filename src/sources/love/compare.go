package love

import (
	"strings"

	"dkfbasel.ch/covid-evidence/stores"
)

// convertRecords will convert the clinicaltrials.gov records to covebasic
func convertRecords(screeningRecords []stores.Record, basicRecords []stores.Record,
	basicIndex stores.Index) ([]*stores.Record, []*stores.Record) {

	sourceName := "LOVE database"

	// initialize the updates/inserts
	add := []*stores.Record{}
	update := []*stores.Record{}

	for _, s := range screeningRecords {

		sourceID := s.Field("url")

		if !strings.Contains(strings.ToLower(s.Field("cove_screening")), "inclu") &&
			!strings.Contains(strings.ToLower(s.Field("cove_screening")), "preprint") {
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
		if alreadyInBasic && !recordFound {
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

			r.Fields["id"] = ""
		}

		r.Update("reviewer_name", s.Field("reviewer_name"), nil)

		r.Update("entry_type", "results_pub", nil)

		r.Update("title", s.Field("title"), nil)

		r.Update("abstract", s.Field("abstract"), nil)

		r.Update("authors", s.Field("author"), nil)

		r.Update("journal", s.Field("publication_title"), nil)

		r.Update("doi", s.Field("doi"), nil)

		r.Update("publication_trial_registration", s.Field("registry"), nil)

		r.Update("status_date", s.Field("date_added"), toIsoDate)

		r.Update("randomized", "randomized", nil)

		// nothing to do, if the record was not changed
		if r.IsUpdated && alreadyInBasic {
			update = append(update, &r)
			continue
		}

		if r.IsUpdated {
			add = append(add, &r)
		}
	}

	return add, update
}
