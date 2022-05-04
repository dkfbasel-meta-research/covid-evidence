package love

import (
	"dkfbasel.ch/covid-evidence/stores"
	"dkfbasel.ch/covid-evidence/stores/mysql"
)

// ImportScreening will import new studies from clinicaltrials gov into the ninox database
func (love *Love) ImportScreening(data []map[string]string) error {

	// // initialize records to import in ninox
	updateDatabase := []*stores.Record{}

	// go through all rows of the csv file (records from love database)
	for _, row := range data {

		// create a new record
		record := stores.Record{}

		// initialize fields
		record.Fields = make(map[string]interface{})

		// add all fields to the record
		for databaseColumnName := range mysql.LoveDict {
			record.Fields[databaseColumnName] = ""
		}

		// add all values to the record
		for key, value := range row {
			record.Fields[key] = value
		}

		updateDatabase = append(updateDatabase, &record)
	}

	err := love.Store.Update("love", updateDatabase)
	if err != nil {
		return err
	}

	return nil
}
