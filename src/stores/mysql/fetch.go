package mysql

import (
	"strconv"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// Fetch ...
func (s Store) Fetch(stmt string) ([]stores.Record, error) {

	rows, err := s.DB.Queryx(stmt)
	if err != nil {
		return nil, err
	}

	logger.Info("start reading")
	results, err := s.scanToMap(rows)
	if err != nil {
		return nil, err
	}
	logger.Info("rows", logger.Any("len", len(results)))

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
