package mysql

import (
	"database/sql"

	"dkfbasel.ch/covid-evidence/logger"
)

// SaveIntervention if not already exist
func (s *Store) SaveIntervention(coveID int32, interventions []map[string]string) error {

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM cove_basic_interventions WHERE cove_basic_id = ?;`, coveID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, intervention := range interventions {
		var interventionID int32
		// TODO: "SELECT id FROM interventions WHERE lower(?) LIKE CONCAT('%', name, '%')"
		err = tx.Get(&interventionID, `
			SELECT id FROM interventions WHERE name = ? LIMIT 1;
		`, intervention["name"])
		if err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}

		if err != nil {
			if !s.DryRun {
				_, err := tx.Exec(`
					INSERT INTO interventions (type, name) VALUES (?, ?);
				`, intervention["type"], intervention["name"])
				if err != nil {
					tx.Rollback()
					return err
				}
				err = tx.Get(&interventionID, `
					SELECT LAST_INSERT_ID();
				`)
				if err != nil {
					tx.Rollback()
					return err
				}
			}

		}

		if !s.DryRun {
			_, err = tx.Exec(`
				INSERT INTO cove_basic_interventions (cove_basic_id, interventions_id) 
				SELECT ?, ? FROM DUAL 
				WHERE NOT EXISTS (
					SELECT * FROM cove_basic_interventions 
					WHERE cove_basic_id = ? AND interventions_id = ? LIMIT 1
				);`, coveID, interventionID, coveID, interventionID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// FindExistingInterventions that and save it
func (s *Store) FindExistingInterventions(coveID int32, intervention string) error {
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	var interventionIDs []int32
	err = tx.Select(&interventionIDs, `
		SELECT id FROM interventions WHERE lower(?) LIKE CONCAT('%', name, '%');
	`, intervention)
	if err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return logger.NewError("could not found interventions", err)
		}
		if err == sql.ErrNoRows {
			logger.Info("No results")
			tx.Rollback()
			return nil
		}
	}
	logger.Info("Results")

	logger.Info("found: ", logger.Any("intervention_id", len(interventionIDs)), logger.Any("cove_id", coveID))

	if !s.DryRun {
		for _, interventionID := range interventionIDs {
			_, err = tx.Exec(`
				INSERT INTO cove_basic_interventions (cove_basic_id, interventions_id) 
				SELECT ?, ? FROM DUAL 
				WHERE NOT EXISTS (
					SELECT * FROM cove_basic_interventions 
					WHERE cove_basic_id = ? AND interventions_id = ? LIMIT 1
				);`, coveID, interventionID, coveID, interventionID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
