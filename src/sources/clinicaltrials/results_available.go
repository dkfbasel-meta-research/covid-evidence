package clinicaltrials

import (
	"strings"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// resultsAvailable will check in the registry and in the database if results are published
func resultsAvailable(store stores.IStore) func(string) (interface{}, bool) {

	return func(m string) (interface{}, bool) {
		fields := strings.Split(m, "$$")

		// get source id from message
		sourceID := fields[0]

		// get results_date from message
		resultsDate := fields[1]
		hasRegistryResults := resultsDate != ""

		hasResults, err := store.CheckForResults(sourceID)
		if err != nil {
			logger.NewError("could not check for results", err)
			hasResults = false
		}

		if hasRegistryResults && hasResults {
			return `["Publication","Registry"]`, true
		}

		if hasRegistryResults {
			return `["Registry"]`, true
		}

		if hasResults {
			return `["Publication"]`, true
		}

		return `["No"]`, true
	}
}
