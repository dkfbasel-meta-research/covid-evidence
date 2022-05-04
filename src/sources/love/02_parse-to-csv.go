package love

import (
	"encoding/csv"
	"io"
	"strings"
)

// Parse have nothing to do, because import is done manually
func (love *Love) Parse(data string) ([]map[string]string, error) {

	r := strings.NewReader(data)
	csvReader := csv.NewReader(r)

	isTitle := true
	var title []string
	content := [][]string{}
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if isTitle {
			title = row
			isTitle = false
			continue
		}

		content = append(content, row)
	}

	csvData := make([]map[string]string, len(content))
	for i, study := range content {
		row := make(map[string]string)

		for k, key := range title {
			row[key] = study[k]
		}

		csvData[i] = row
	}

	return csvData, nil
}
