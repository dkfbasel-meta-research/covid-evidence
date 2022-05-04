package mysql

import "dkfbasel.ch/covid-evidence/stores"

// FetchScreening ...
func (s *Store) FetchScreening(source string, fields []string, filter string) ([]stores.Record, error) {
	if source == "clinicaltrials.gov" {
		return s.FetchRecords("screening_clinicaltrials", fields, filter)
	}
	if source == "ictrp" {
		return s.FetchRecords("screening_ictrp", fields, filter)
	}
	if source == "love" {
		return s.FetchRecords("screening_love", fields, filter)
	}
	if source == "johns-hopkins" {
		return s.FetchRecords("johnshopkins", fields, filter)
	}
	return nil, nil
}
