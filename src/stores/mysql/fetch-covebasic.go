package mysql

import (
	"fmt"
	"strconv"
	"strings"

	"dkfbasel.ch/covid-evidence/stores"
)

const CoveBasicTable = "covebasic"
const CoveBasicExlusionsTable = "exclusions"

// fetchCoveBasic will fetch all records from covebasic and return an index
// for source_ids and information if the given id is contained in covebasic/covebasic
// or covebasic/exclusions
func (s *Store) fetchCoveBasic(sources ...string) (records []stores.Record, index stores.Index, err error) {

	// fetch all records from covebasic/covebasic
	basicIncluded, err := s.FetchRecords("cove_basic", nil, "")
	if err != nil {
		return nil, nil,
			fmt.Errorf("could not fetch covebasic records from mysql: %w", err)
	}

	// index of items in covebasic or covebasic/exclusions
	index = make(stores.Index)

	// contains will check if the source is part of the sources list
	contains := func(list []string, source string) bool {

		// convert source to lowercase for comparison
		source = strings.ToLower(source)

		// return true if no comparison list is provided
		if list == nil {
			return true
		}
		for i := range list {
			if list[i] == source {
				return true
			}
		}
		return false
	}

	// convert sources to lowercase for comparison
	for i := range sources {
		sources[i] = strings.ToLower(sources[i])
	}

	// // check exluded first, so that inclusion takes precedence over exlusions
	// for i, record := range basicExcluded {

	// 	id := record.Field("source_id")
	// 	source := record.Field("source")

	// 	if !contains(sources, source) {
	// 		continue
	// 	}
	// 	index.Set(id, record.ID, CoveBasicExlusionsTable, &basicExcluded[i])
	// }

	for i, record := range basicIncluded {

		id := record.Field("source_id")
		source := strings.ToLower(record.Field("source"))

		if !contains(sources, source) {
			continue
		}

		basicID, err := strconv.ParseInt(record.Field("id"), 10, 32)
		if err == nil {
			basicIncluded[i].BasicID = int32(basicID)
		}

		// overwrite items contained in exlusions
		info, ok := index.Get(id)
		if ok && info.Table == CoveBasicTable {
			continue
		}

		index.Set(id, record.ID, CoveBasicTable, &basicIncluded[i])
	}

	// return the records and the corresponding index
	return basicIncluded, index, nil

}
