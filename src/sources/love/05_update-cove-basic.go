package love

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// UpdateCoveBasic will transfer the records from the screening table to the covebasic table
func (love *Love) UpdateCoveBasic() error {

	screeningRecords, err := love.Store.FetchScreening("love", nil, "cove_screening LIKE '%incl%' OR cove_screening LIKE '%preprint%'")
	if err != nil {
		return logger.NewError("could not fetch screening", err)
	}

	coveBasicRecords, basicIndex, err := love.Store.FetchBasic("LOVE database")
	if err != nil {
		return logger.NewError("could not fetch cove basic table entries", err)
	}

	logger.Info("screening records", logger.Any("len", len(screeningRecords)))
	logger.Info("basic records", logger.Any("len", len(coveBasicRecords)))
	logger.Info("index records", logger.Any("len", len(basicIndex)))

	// convert the export to our basic table
	add, update := convertRecords(
		screeningRecords,
		coveBasicRecords,
		basicIndex,
	)

	logger.Info("updated records", logger.Any("len", len(update)))
	logger.Info("new records", logger.Any("len", len(add)))

	// import the new basic records into the database
	err = love.Store.AddBasic(add)
	if err != nil {
		return err
	}

	// update records in cove basic
	err = love.Store.UpdateBasic(update)
	if err != nil {
		return err
	}

	logger.Info("update of covebasic completed")

	return nil

}
