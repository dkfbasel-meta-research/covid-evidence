package clinicaltrials

import (
	"encoding/json"
	"fmt"
	"dkfbasel.ch/covid-evidence/logger"
	"strings"

	"github.com/tidwall/gjson"
)

// Parse will convert the given information to the specified data model
func (ct *ClinicalTrials) Parse(data string) ([]map[string]string, error) {

	// parse the content of the json file
	var studies []json.RawMessage
	err := json.Unmarshal([]byte(data), &studies)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	logger.Info("number of studies", logger.Any("len", len(studies)))

	csvData := make([]map[string]string, len(studies))

	// iterate through all studies in the dataset
	for i, study := range studies {

		// initialize a new study information map
		row := make(map[string]string)

		// initialize a searchable json structure
		study := gjson.GetBytes(study, "Study")

		// try to extract the field content according to the field map
		for _, field := range fieldMap {

			// skip all fields without a search path specified
			if field.Search == "" {
				continue
			}

			value := study.Get(field.Search)
			if !value.Exists() {
				// logger.Info("could not find field", logger.Any("field", field.Search))
				continue
			}

			// concatenate arrays with semicolon
			// note: might be better to keep items separated in the future
			if value.IsArray() {
				values := value.Array()
				asStr := make([]string, len(values))
				for j, item := range values {
					asStr[j] = item.String()
				}
				row[field.Name] = strings.Join(asStr, "; ")
				continue
			}

			row[field.Name] = value.String()
		}

		csvData[i] = row
	}

	return csvData, nil
}
