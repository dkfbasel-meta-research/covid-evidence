package johnshopkins

import (
	"strconv"
	"time"

	"dkfbasel.ch/covid-evidence/logger"
)

// createRecord will uniform Covid19 cases data from the johns hopkins university to
// import to the database
func createRecord(countryString, dateString, casesString string) (map[string]string, error) {
	record := make(map[string]string)

	date, err := time.Parse("1/2/06", dateString)
	if err != nil {
		return nil, logger.NewError("could not parse date", err)
	}

	week, year := date.ISOWeek()

	record["country"] = countryString
	record["week"] = strconv.Itoa(week)
	record["year"] = strconv.Itoa(year)
	record["date"] = date.Format("2006-02-01")
	record["cases"] = casesString

	return record, nil
}

// Clean will check the number of new entries in the source
func (jh *JohnsHopkins) Clean(data []map[string]string) ([]map[string]string, int, int, error) {

	var lastData map[string]string
	currentData := make(map[string]string)
	records := make([]map[string]string, 0)
	lastCountry := ""
	currentCountry := ""
	regions := false

	for i := range data {
		lastData = currentData
		lastCountry = currentCountry
		currentData = make(map[string]string)
		for k, v := range data[i] {
			if v == "" {
				continue
			}
			if k == "Country/Region" {
				currentCountry = v
			} else if k == "Province/State" && v != "" {
				regions = true
			} else if k == "Lat" || k == "Long" {
				continue
			} else {
				currentData[k] = v
			}
		}

		if lastCountry == currentCountry && regions {
			for date, currentValue := range currentData {
				oldValue, ok := lastData[date]
				if ok {
					if oldValue == "" {
						oldValue = "0"
					}
					if currentValue == "" {
						currentValue = "0"
					}
					oldCases, _ := strconv.ParseInt(oldValue, 10, 64)
					currentCases, _ := strconv.ParseInt(currentValue, 10, 64)
					newCases := oldCases + currentCases
					currentData[date] = strconv.FormatInt(newCases, 10)
				} else {
					currentData[date] = currentValue
				}
			}
		}

		if lastCountry != currentCountry {
			for currentDateKey, currentCases := range lastData {

				record, err := createRecord(lastCountry, currentDateKey, currentCases)
				if err != nil {
					logger.Error("could not create record", err, logger.String("country", lastCountry),
						logger.String("cases", currentCases))
				}

				records = append(records, record)
			}
		}
	}

	// add the last country in the list that is not covered by the the loop above
	for currentDateKey, currentCases := range lastData {
		record, err := createRecord(lastCountry, currentDateKey, currentCases)
		if err != nil {
			logger.Error("could not create record", err, logger.String("country", lastCountry),
				logger.String("cases", currentCases))
		}

		records = append(records, record)
	}

	return records, len(records), 0, nil
}
