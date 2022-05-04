package mysql

import (
	"fmt"
	"strings"
)

// CleanTrials ...
func (s *Store) CleanTrials(trialIDs []string) ([]string, error) {

	stmt := fmt.Sprintf(`
		SELECT id FROM trials
		WHERE love_id IS NOT NULL AND love_id <> ''
		AND love_id NOT IN ('%s');`, strings.Join(trialIDs, "','"))

	trialsToDelete := []int64{}
	err := s.DB.Select(&trialsToDelete, stmt)
	if err != nil {
		return nil, err
	}

	if len(trialsToDelete) > 0 {
		// delete trials
		deleteTrials := strings.Trim(strings.Replace(fmt.Sprint(trialsToDelete), " ", ",", -1), "[]")
		stmt = fmt.Sprintf(`
			DELETE trials, registry_cove_basic, trials_cove_basic
			FROM trials
			LEFT JOIN registry_cove_basic ON registry_cove_basic.trials_id = trials.id
			LEFT JOIN trials_cove_basic ON trials_cove_basic.trials_id = trials.id
			WHERE trials.id IN (%s);`, deleteTrials)

		_, err = s.DB.Exec(stmt)
		if err != nil {
			return nil, err
		}
	}

	stmt = fmt.Sprintf(`
		SELECT love_id FROM trials
		WHERE love_id IS NOT NULL AND love_id <> ''
		AND love_id IN ('%s');`, strings.Join(trialIDs, "','"))
	alreadyInTrials := []string{}
	err = s.DB.Select(&alreadyInTrials, stmt)
	if err != nil {
		return nil, err
	}

	return alreadyInTrials, nil
}
