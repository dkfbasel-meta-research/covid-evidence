package johnshopkins

import (
	"strconv"
	"time"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// ImportScreening will import new studies from clinicaltrials gov into the ninox database
func (jh *JohnsHopkins) ImportScreening(data []map[string]string) error {

	// sum up covid cases by week
	// map[year][week][country] >> cases

	weeklyCases := make(map[int]map[int]map[string]int)

	// go through all rows of the csv file (records from ct.gov)
	for _, row := range data {

		country := ""
		year := 0
		week := 0
		cases := 0

		// add all fields to the record
		for key, value := range row {
			if key == "" {
				continue
			}

			if key == "date" {
				date, err := time.Parse("1/2/06", value)
				if err != nil {
					logger.Info("could not parse date", logger.String("date_string", value))
					continue
				}
				year, week = date.ISOWeek()

				continue
			}

			if key == "country" {
				country = value

				continue
			}

			if key == "cases" {
				weekCases, err := strconv.Atoi(value)
				if err != nil {
					logger.Info("could not parse cases", logger.String("case_string", value))
					continue
				}
				cases = weekCases
			}
		}

		// check if map exists for current year
		if _, ok := weeklyCases[year]; !ok {
			weeklyCases[year] = make(map[int]map[string]int)
		}
		if _, ok := weeklyCases[year][week]; !ok {
			weeklyCases[year][week] = make(map[string]int)
		}
		if currentCases, ok := weeklyCases[year][week][country]; !ok {
			weeklyCases[year][week][country] = cases
		} else if currentCases < cases {
			weeklyCases[year][week][country] = cases
		}
	}

	// initialize records to import to the database
	updateDatabase := []*stores.Record{}

	for year, yearMap := range weeklyCases {
		for week, weekMap := range yearMap {
			for country, cases := range weekMap {
				record := stores.Record{}
				record.Fields = make(map[string]interface{})
				record.Fields["year"] = year
				record.Fields["week"] = week
				record.Fields["country"] = country
				record.Fields["cases"] = cases

				// create records for each country and week
				updateDatabase = append(updateDatabase, &record)
			}
		}
	}

	err := jh.Store.Update("johns-hopkins", updateDatabase)
	if err != nil {
		return err
	}

	return nil
}
