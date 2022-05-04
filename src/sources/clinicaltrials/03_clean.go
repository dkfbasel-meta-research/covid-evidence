package clinicaltrials

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// Clean will check the number of new entries in the source
func (ct *ClinicalTrials) Clean(data []map[string]string) ([]map[string]string, int, int, error) {

	screening, err := ct.Store.FetchScreening("clinicaltrials.gov", []string{"nct_id"}, "")
	if err != nil {
		return nil, 0, 0, logger.NewError("could not fetch screening", err)
	}

	// map field names to database field names
	fieldInputRowIndex := make(map[string]string)
	for _, field := range fieldMap {
		if field.Ninox != "" {
			fieldInputRowIndex[field.Ninox] = field.Name
		}
	}

	mappedData := make([]map[string]string, len(data))
	for i := range data {
		mappedData[i] = make(map[string]string)
		for newKey, oldKey := range fieldInputRowIndex {
			mappedData[i][newKey] = data[i][oldKey]
		}
	}
	data = mappedData

	// numbers
	duplicatesInSource := 0
	entriesToAdd := 0

	// check if there are duplicates in source
	sourceIndex := make(map[string]int)
	screeningIndex := make(map[string]string)

	// create the sourceIndex
	for i := range data {
		sourceID := data[i]["nct_id"]
		if _, ok := sourceIndex[sourceID]; ok {
			sourceIndex[sourceID] = sourceIndex[sourceID] + 1
		} else {
			sourceIndex[sourceID] = 1
		}
	}

	// create a screening index
	for i := range screening {
		sourceID := screening[i].Field("nct_id")
		if _, ok := screeningIndex[sourceID]; !ok {
			screeningIndex[sourceID] = screening[i].Field("id")
		}
	}

	for i := range data {
		sourceID := data[i]["nct_id"]
		if sourceIndex[sourceID] > 1 {
			data[i]["duplicates"] = "1"
			duplicatesInSource++
		} else {
			data[i]["duplicates"] = "0"
		}
		if id, ok := screeningIndex[sourceID]; ok {
			data[i]["id"] = id
		} else {
			data[i]["id"] = ""
			entriesToAdd++
		}
	}

	return data, entriesToAdd, len(screening), nil
}
