package pipeline

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"dkfbasel.ch/covid-evidence/logger"
)

// StartSourceStep ...
func (p *Pipeline) StartSourceStep() string {
	workflowID := fmt.Sprintf("%d", time.Now().Unix())
	exportPath := filepath.Join(p.Tmp, workflowID)
	err := os.MkdirAll(exportPath, 0777)

	logger.Info("Start pipeline")

	if err != nil {
		logger.Error("could not find export path", err)
		return ""
	}

	for _, source := range p.Sources {

		logger.Info("Start", logger.String("source", source.GetID()))

		sourcePath := filepath.Join(exportPath, source.GetID())
		err := os.MkdirAll(sourcePath, 0777)
		if err != nil {
			logger.Error("could not find source path", err)
			continue
		}

		sourceString, err := source.Fetch()
		if err != nil {
			logger.Error("could not fetch source information", err)
			continue
		}
		if sourceString == "" {
			logger.Info("source data is empty", logger.String("source", source.GetID()))
			continue
		}

		err = ioutil.WriteFile(filepath.Join(sourcePath, "source.txt"), []byte(sourceString), 0666)
		if err != nil {
			logger.Error("could not write data file", err)
			continue
		}

		sourceMap, err := source.Parse(sourceString)
		if err != nil {
			logger.Error("could not parse source data", err)
			continue
		}
		err = writeCSV(filepath.Join(sourcePath, "source.csv"), sourceMap)
		if err != nil {
			logger.Error("could not write CSV", err)
			continue
		}

		data, _, _, err := source.Clean(sourceMap)
		if err != nil {
			logger.Error("could not clean data", err)
			continue
		}
		err = writeCSV(filepath.Join(sourcePath, "data.csv"), data)
		if err != nil {
			logger.Error("could write data", err)
			continue
		}
	}
	return workflowID
}
