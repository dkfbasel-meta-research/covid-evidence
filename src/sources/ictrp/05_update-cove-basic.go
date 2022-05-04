package ictrp

import (
	"strings"

	"dkfbasel.ch/covid-evidence/logger"

	"dkfbasel.ch/covid-evidence/stores"
)

// UpdateCoveBasic will transfer the records from the screening table to the covebasic table
func (who *ICTRP) UpdateCoveBasic() error {

	screeningRecords, err := who.Store.FetchScreening("ictrp", nil, "((cove_screening = '' OR cove_screening IS NULL) AND study_type LIKE '%interven%') OR cove_screening LIKE '%incl%'")
	// screeningRecords, err := who.Store.FetchScreening("ictrp", nil, `
	// 	cove_screening LIKE '%include%'
	// `)
	if err != nil {
		return logger.NewError("could not fetch screening", err)
	}

	coveBasicRecords, basicIndex, err := who.Store.FetchBasic("ictrp", "clinicaltrials.gov")
	if err != nil {
		return logger.NewError("could not fetch cove basic table entries", err)
	}

	filteredScreening := []stores.Record{}
	for _, record := range screeningRecords {
		studyType := record.Field("study_type")
		if strings.Contains(strings.ToLower(studyType), "intervention") {
			filteredScreening = append(filteredScreening, record)
		}
	}
	screeningRecords = filteredScreening

	logger.Info("fetched screening records", logger.Any("len", len(screeningRecords)))
	logger.Info("fetched basic records", logger.Any("len", len(coveBasicRecords)))
	logger.Info("index contains records", logger.Any("len", len(basicIndex)))

	// convert the export to our basic table
	add, update := convertRecords(
		who.Store,
		screeningRecords,
		coveBasicRecords,
		basicIndex,
	)

	logger.Info("updated records", logger.Any("len", len(update)))
	logger.Info("new records", logger.Any("len", len(add)))

	// import the new basic records into ninox
	err = who.Store.AddBasic(add)
	if err != nil {
		return err
	}

	// update the basic records in the database
	err = who.Store.UpdateBasic(update)
	if err != nil {
		return err
	}

	logger.Info("update of covebasic completed")

	return nil

}
