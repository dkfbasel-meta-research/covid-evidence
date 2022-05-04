package clinicaltrials

import "dkfbasel.ch/covid-evidence/stores"

const requestURL = "https://www.clinicaltrials.gov/api/query/full_studies"
const query = "(wuhan AND (coronavirus OR corona virus OR pneumonia virus)) OR COVID19 OR COVID-19 OR COVID 19 OR coronavirus 2019 OR corona virus 2019 OR SARS-CoV-2 OR SARSCoV2 OR SARS2 OR SARS-2 OR 2019 nCoV OR ((novel coronavirus OR novel corona virus) AND 2019)"

// ClinicalTrials ...
type ClinicalTrials struct {
	ID    string
	Store stores.IStore
}

// NewSource initialize a new ictrp source
func NewSource(store stores.IStore) *ClinicalTrials {
	source := ClinicalTrials{}
	source.ID = "clinicaltrials"
	source.Store = store
	return &source
}

// GetID will return the store identifier
func (ct *ClinicalTrials) GetID() string {
	return ct.ID
}
