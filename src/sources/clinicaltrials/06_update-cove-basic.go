package clinicaltrials

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// UpdateCoveBasic will transfer the records from the screening table to the covebasic table
func (ct *ClinicalTrials) UpdateCoveBasic() error {

	screening, err := ct.Store.Fetch(`
		SELECT screening_clinicaltrials.* FROM screening_clinicaltrials
		WHERE (
			(
				cove_screening = '' OR cove_screening IS NULL
			) AND study_type LIKE '%interven%'
		) OR (cove_screening LIKE '%inclu%');
	`)
	// screening, err := ct.Store.Fetch(`
	// 	SELECT screening_clinicaltrials.* FROM screening_clinicaltrials
	// 	WHERE cove_screening LIKE '%include%';
	// `)
	if err != nil {
		return logger.NewError("could not fetch screening", err)
	}

	basic, basicIndex, err := ct.Store.FetchBasic("ictrp", "clinicaltrials.gov")
	if err != nil {
		return logger.NewError("could not fetch cove basic table entries", err)
	}

	logger.Info("screening records", logger.Any("len", len(screening)))
	logger.Info("basic records", logger.Any("len", len(basic)))
	logger.Info("index records", logger.Any("len", len(basicIndex)))

	// convert the export to our basic table
	add, update := convertRecords(
		ct.Store,
		screening,
		basic,
		basicIndex)

	// fmt.Println(len(add))
	// fmt.Println(len(update))
	// os.Exit(1)

	logger.Info("updated records", logger.Any("len", len(update)))
	logger.Info("new records", logger.Any("len", len(add)))

	// import the new basic records into ninox
	err = ct.Store.AddBasic(add)
	if err != nil {
		return logger.NewError("could not add new entries to cove basic", err)
	}

	// import the new basic records into ninox
	err = ct.Store.UpdateBasic(update)
	if err != nil {
		return logger.NewError("could not update cove basic", err)
	}

	logger.Info("update of covebasic completed")

	return nil
}
