package mysql

import (
	"fmt"
	"strconv"
	"strings"

	"dkfbasel.ch/covid-evidence/stores"
)

// FetchRecords will fetch all record from the given mysql table
func (s *Store) FetchRecords(table string, fields []string, filters string) ([]stores.Record, error) {

	fieldSelection := "*"
	if fields != nil && len(fields) != 0 {
		fields = append(fields, "id")
		fieldSelection = strings.Join(fields, ", ")
	}

	stmt := fmt.Sprintf(`SELECT %s FROM %s`, fieldSelection, table)
	if filters != "" {
		stmt = stmt + fmt.Sprintf(" WHERE %s", filters)
	}

	rows, err := s.DB.Queryx(stmt)
	if err != nil {
		return nil, err
	}
	results, err := s.scanToMap(rows)
	if err != nil {
		return nil, err
	}

	rows.Close()

	records := []stores.Record{}

	for i := range results {
		record := stores.Record{}
		if v, ok := results[i]["id"]; ok {
			recordID := v.([]byte)
			recordIDString := string(recordID)
			record.ID, err = strconv.Atoi(recordIDString)
			if err != nil {
				return nil, err
			}
			record.Fields = results[i]
			records = append(records, record)
		} else {
			continue
		}
	}

	return records, nil
}
