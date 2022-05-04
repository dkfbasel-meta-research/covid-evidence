package pipeline

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// StartCoveBasicStep ...
func (p *Pipeline) StartCoveBasicStep() {

	logger.Info("start cove basic import")

	message := make(map[string]map[string]int)

	for _, source := range p.Sources {

		message[source.GetID()] = make(map[string]int)
		logger.Info("start", logger.String("source", source.GetID()))

		err := source.UpdateCoveBasic()
		if err != nil {
			logger.Error("could not update cove basic", err)
			continue
		}

		err = source.UpdateScreening()
		if err != nil {
			logger.Error("could not update cove screening", err)
			continue
		}
	}
}
