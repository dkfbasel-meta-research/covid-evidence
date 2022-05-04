package mysql

import (
	"fmt"
	"strings"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// SetTopics ...
func (s *Store) SetTopics(topics []stores.Record) error {

	for _, r := range topics {
		id := r.Field("id")
		columns := strings.Split(strings.TrimSpace(r.Field("keyword_search_columns")), ",")
		keywords := strings.Split(strings.TrimSpace(r.Field("keywords")), ",")
		name := strings.TrimSpace(r.Field("name"))

		if len(columns) == 0 || len(keywords) == 0 {
			return logger.NewError("no column or keyword specified", nil)
		}

		whereClauses := []string{}
		for _, c := range columns {
			for _, k := range keywords {
				whereClauses = append(whereClauses, fmt.Sprintf(`cove_basic.%s LIKE '%%%s%%'`, c, k))
			}
		}

		stmt := fmt.Sprintf(`
			INSERT INTO topic_cove_basic (cove_basic_id, topic_id)
			SELECT cove_basic.id, %s FROM cove_basic
			WHERE is_rct='yes' and is_covid='yes' and
				is_duplicate='no' AND is_trial='yes' AND is_observational = 'no'
			AND (%s)
			AND cove_basic.id NOT IN (
				SELECT cove_basic_id FROM topic_cove_basic
				WHERE topic_cove_basic.topic_id = %s
				AND cove_basic.id IS NOT NULL
			);`, id, strings.Join(whereClauses, " OR "), id)

		if !s.DryRun {
			_, err := s.DB.Exec(stmt)
			if err != nil {
				logger.Error("could not set topic", err, logger.String("topic", name))
			} else {
				logger.Info("topic is set", logger.String("topic", name))
			}
		}
	}

	return nil
}
