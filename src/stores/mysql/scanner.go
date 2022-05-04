package mysql

import (
	"fmt"

	"dkfbasel.ch/covid-evidence/logger"
	"github.com/jmoiron/sqlx"
)

func (s *Store) scanToMap(rows *sqlx.Rows) ([]map[string]interface{}, error) {

	results := make([]*map[string]interface{}, 0)

	index := 0
	for rows.Next() {

		index = index + 1
		if index%1000 == 0 {
			logger.Info(fmt.Sprintf("%d", index))
		}

		result := make(map[string]interface{})

		err := rows.MapScan(result)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	resultsMapped := make([]map[string]interface{}, 0)
	for i := range results {
		resultsMapped = append(resultsMapped, *results[i])
	}

	return resultsMapped, nil
}
