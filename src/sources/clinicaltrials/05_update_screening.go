package clinicaltrials

// UpdateScreening will set the cove_screening variable in the database
func (ct *ClinicalTrials) UpdateScreening() error {

	return ct.Store.SetScreening("clinicaltrials.gov")
}
