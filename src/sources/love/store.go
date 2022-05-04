package love

import (
	"dkfbasel.ch/covid-evidence/stores"
)

// FetchScreening will fetch teh screening records
func (love *Love) FetchScreening(store stores.IStore, filter string) ([]stores.Record, error) {
	return store.FetchScreening("ictrp", nil, "")
}
