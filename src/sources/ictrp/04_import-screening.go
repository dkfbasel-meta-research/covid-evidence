package ictrp

import (
	"dkfbasel.ch/covid-evidence/stores"
)

// ImportScreening will import new studies from clinicaltrials gov into the ninox database
func (who *ICTRP) ImportScreening(data []map[string]string) error {

	// initialize records to import in ninox
	updateDatabase := []*stores.Record{}

	// go through all rows of the csv file (records from ctgov)
	for _, row := range data {

		// create a new record
		record := stores.Record{}

		// initialize fields
		record.Fields = make(map[string]interface{})

		// add all fields to the record
		for key, value := range row {
			record.Fields[key] = value
		}

		updateDatabase = append(updateDatabase, &record)
	}

	err := who.Store.Update("ictrp", updateDatabase)
	if err != nil {
		return err
	}

	return nil
}
