package mysql

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// SetScreening ...
func (s *Store) SetScreening(source string) error {

	var stmt1 string
	var stmt2 string
	var stmt3 string
	var stmt4 string
	var stmt5 string
	var stmt6 string

	if source == "clinicaltrials.gov" {
		stmt1 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not RCT'
			WHERE cove_screening <> 'automatic exclusion'
			AND nct_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
				AND is_rct = 'no'
			);`

		stmt2 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not covid'
			WHERE cove_screening <> 'automatic exclusion'
			AND nct_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
				AND is_covid = 'no'
			);`

		stmt3 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not interventional'
			WHERE cove_screening <> 'automatic exclusion'
			AND nct_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
				AND is_observational = 'yes'
			);`

		stmt4 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not a trial'
			WHERE cove_screening <> 'automatic exclusion'
			AND nct_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
				AND is_trial = 'no'
			);`

		stmt5 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic inclusion'
			WHERE (
				screening_clinicaltrials.cove_screening IS NULL
				OR screening_clinicaltrials.cove_screening = ''
			) AND nct_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
			);`

		stmt6 = `
			UPDATE screening_clinicaltrials
			SET cove_screening = 'automatic exclusion'
			WHERE cove_screening = '' OR cove_screening IS NULL
			AND nct_id NOT IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'clinicaltrials.gov'
			);`
	} else if source == "ictrp" {
		stmt1 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not RCT'
			WHERE cove_screening <> 'automatic exclusion'
			AND trial_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
				AND is_rct = 'no'
			);`

		stmt2 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not covid'
			WHERE cove_screening <> 'automatic exclusion'
			AND trial_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
				AND is_covid = 'no'
			);`

		stmt3 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not interventional'
			WHERE cove_screening <> 'automatic exclusion'
			AND trial_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
				AND is_observational = 'yes'
			);`

		stmt4 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic exclusion', cove_screening_comment = 'not a trial'
			WHERE cove_screening <> 'automatic exclusion'
			AND trial_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
				AND is_trial = 'no'
			);`

		stmt5 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic inclusion'
			WHERE (
				screening_ictrp.cove_screening IS NULL
				OR screening_ictrp.cove_screening = ''
			)
			AND trial_id IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
			);`

		stmt6 = `
			UPDATE screening_ictrp
			SET cove_screening = 'automatic exclusion'
			WHERE cove_screening = '' OR cove_screening IS NULL
			AND trial_id NOT IN (
				SELECT source_id FROM cove_basic
				WHERE source = 'ICTRP'
			);`
	} else if source == "love" {
		return nil
	}

	if !s.DryRun {
		_, err := s.DB.Exec(stmt1)
		if err != nil {
			return err
		}
	}
	logger.Info("not RCT set")

	if !s.DryRun {
		_, err := s.DB.Exec(stmt2)
		if err != nil {
			return err
		}
	}
	logger.Info("not covid set")

	if !s.DryRun {
		_, err := s.DB.Exec(stmt3)
		if err != nil {
			return err
		}
	}
	logger.Info("not interventional")

	if !s.DryRun {
		_, err := s.DB.Exec(stmt4)
		if err != nil {
			return err
		}
	}
	logger.Info("not a trial")

	if !s.DryRun {
		_, err := s.DB.Exec(stmt5)
		if err != nil {
			return err
		}
	}
	logger.Info("inclusion set")

	if !s.DryRun {
		_, err := s.DB.Exec(stmt6)
		if err != nil {
			return err
		}
	}
	logger.Info("exclusion set")

	return nil
}
