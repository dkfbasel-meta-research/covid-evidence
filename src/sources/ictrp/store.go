package ictrp

import (
	"dkfbasel.ch/covid-evidence/stores"
)

// FetchScreening will fetch the screening records
func (who *ICTRP) FetchScreening(store stores.IStore, filter string) ([]stores.Record, error) {
	return store.FetchScreening("ictrp", nil, "")
}
