package clinicaltrials

import (
	"strconv"

	"dkfbasel.ch/covid-evidence/stores"
)

// ImportScreening will import new studies from clinicaltrials gov into the ninox database
func (ct *ClinicalTrials) ImportScreening(data []map[string]string) error {

	// initialize records to import in ninox
	updateDatabase := []*stores.Record{}

	// go through all rows of the csv file (records from ct.gov)
	for _, row := range data {

		// create a new record
		record := stores.Record{}

		// initialize fields
		record.Fields = make(map[string]interface{})

		// add all fields to the record
		for key, value := range row {
			if key == "" {
				continue
			}

			if key == "enrollment" {
				asInt, _ := strconv.Atoi(value)
				record.Fields[key] = asInt
				continue
			}

			record.Fields[key] = value
		}

		updateDatabase = append(updateDatabase, &record)
	}

	err := ct.Store.Update("clinicaltrials.gov", updateDatabase)
	if err != nil {
		return err
	}

	return nil
}
