package mysql

import (
	"fmt"
	"strconv"

	"dkfbasel.ch/covid-evidence/logger"
)

// AddTrial ...
func (s *Store) AddTrial(loveID string, publicationID string, registryIDs []string) error {

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `
		INSERT INTO trials (automatically_created, love_id)
		VALUES (1, ?);`

	res, err := tx.Exec(stmt, loveID)
	if err != nil {
		return logger.NewError("could not insert trial", err, logger.String("stmt", stmt))
	}
	trialID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// add publication
	stmt = fmt.Sprintf(`
		INSERT INTO trials_cove_basic (trials_id, cove_basic_id)
		SELECT %d, id FROM cove_basic
		WHERE source_id = ?`, trialID)

	_, err = tx.Exec(stmt, publicationID)
	if err != nil {
		return logger.NewError("could not insert publications", err, logger.String("stmt", stmt))
	}

	// add registry entries
	for _, registryID := range registryIDs {
		stmt = fmt.Sprintf(`
		INSERT INTO registry_cove_basic (trials_id, cove_basic_id)
		SELECT %s, id FROM cove_basic
		WHERE source_id LIKE '%s%s%s'`, strconv.FormatInt(trialID, 10), "%", registryID, "%")
		_, err = tx.Exec(stmt)
		if err != nil {
			return logger.NewError("could not insert registry entries", err, logger.String("stmt", stmt))
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
