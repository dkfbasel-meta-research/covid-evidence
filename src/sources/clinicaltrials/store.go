package clinicaltrials

import (
	"dkfbasel.ch/covid-evidence/stores"
)

// FetchScreening will fetch teh screening records
func (ct *ClinicalTrials) FetchScreening(store stores.IStore, fields []string, filter string) ([]stores.Record, error) {
	return store.FetchScreening("clinicaltrials.org", fields, filter)
}
