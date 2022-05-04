package love

import (
	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores/mysql"
)

// Clean will check the number of new entries in the source
func (love *Love) Clean(data []map[string]string) ([]map[string]string, int, int, error) {

	screening, err := love.Store.FetchScreening("love", []string{"url"}, "")
	if err != nil {
		return nil, 0, 0, logger.NewError("could not fetch screening", err)
	}

	mappedData := make([]map[string]string, len(data))
	for i := range data {
		mappedData[i] = make(map[string]string)
		for key := range data[i] {
			for databaseColumnName, csvColumnName := range mysql.LoveDict {
				if csvColumnName == key {
					mappedData[i][databaseColumnName] = data[i][key]
				}
			}
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
		sourceID := data[i]["url"]
		if _, ok := sourceIndex[sourceID]; ok {
			sourceIndex[sourceID]++
		} else {
			sourceIndex[sourceID] = 1
		}
	}

	// create a screening index
	for i := range screening {
		sourceID := screening[i].Field("url")
		if _, ok := screeningIndex[sourceID]; !ok {
			screeningIndex[sourceID] = screening[i].Field("id")
		}
	}

	for i := range data {
		sourceID := data[i]["url"]
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
