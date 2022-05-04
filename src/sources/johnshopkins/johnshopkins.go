package johnshopkins

import "dkfbasel.ch/covid-evidence/stores"

const requestURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"

// JohnsHopkins ...
type JohnsHopkins struct {
	ID    string
	Store stores.IStore
}

// NewSource initialize a new ictrp source
func NewSource(store stores.IStore) *JohnsHopkins {
	source := JohnsHopkins{}
	source.ID = "johnshopkins"
	source.Store = store
	return &source
}

// GetID will return the store identifier
func (ct *JohnsHopkins) GetID() string {
	return ct.ID
}
