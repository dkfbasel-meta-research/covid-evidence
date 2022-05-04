package pipeline

import (
	"os"
	"path/filepath"

	"dkfbasel.ch/covid-evidence/logger"
)

// StartScreeningStep ...
func (p *Pipeline) StartScreeningStep(workflowID string) {

	// workflowID = "1602598765"

	exportPath := filepath.Join(p.Tmp, workflowID)
	err := os.MkdirAll(exportPath, 0777)
	if err != nil {
		logger.Error("could not find export path", err)
		return
	}

	logger.Info("start screening import")

	message := make(map[string]map[string]int)

	for _, source := range p.Sources {

		message[source.GetID()] = make(map[string]int)
		logger.Info("start", logger.String("source", source.GetID()))

		sourcePath := filepath.Join(exportPath, source.GetID())

		sourceData, err := readCSV(filepath.Join(sourcePath, "data.csv"))
		if err != nil {
			logger.Error("could not read source data", err)
			return
		}

		err = source.ImportScreening(sourceData)
		if err != nil {
			logger.Error("could not import screening", err)
			return
		}
	}
}
