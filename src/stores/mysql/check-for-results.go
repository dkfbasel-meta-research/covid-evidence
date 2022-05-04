package mysql

// CheckForResults will check if a linked entry exists for a given registry id
func (s Store) CheckForResults(id string) (bool, error) {

	var resultsAvailable bool
	stmt := `
		SELECT count(results.source_id) > 0 AS available FROM cove_basic AS results
		INNER JOIN trials_cove_basic ON results.id = trials_cove_basic.cove_basic_id
		INNER JOIN trials ON trials_cove_basic.trials_id = trials.id
		INNER JOIN registry_cove_basic ON registry_cove_basic.trials_id = trials.id
		INNER JOIN cove_basic AS registry ON registry.id = registry_cove_basic.cove_basic_id
		WHERE registry.source_id = ?;`

	err := s.DB.Get(&resultsAvailable, stmt, id)
	if err != nil {
		return false, err
	}

	return resultsAvailable, nil
}
