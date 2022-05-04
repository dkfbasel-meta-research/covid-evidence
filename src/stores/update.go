package stores

import (
	"fmt"
)

const certaintyPrefilled = "prefilled"
const certaintyGenerated = "generated"
const certaintyHuman = "human"
const certaintyVerified = "verified"

var index int = 1

// Update ...
func (r *Record) Update(fieldName string, sourceField interface{}, fn handlerFunc) {

	// define the name of the associated certainty field
	fieldNameCertainty := fmt.Sprintf("%s_certainty", fieldName)

	sourceValue := asString(sourceField)
	currentValue := asString(r.Fields[fieldName])
	currentCertainty := asString(r.Fields[fieldNameCertainty])

	// is the content generated
	isGenerated := false

	// use a custom handler function for the variable if specified
	if fn != nil {
		sourceField, isGenerated = fn(sourceValue)
		sourceValue = asString(sourceField)
	}

	// skip the update procedure if the value hasn't change
	if currentValue == sourceValue {
		return
	}

	// skip the update if the new value is empty
	if sourceValue == "" {
		return
	}

	// update the following fields independent of their certainty value:
	// n_enrollment, status, status_date
	if fieldName == "n_enrollment" || fieldName == "status" || fieldName == "status_date" {
		r.IsUpdated = true
		r.Fields[fieldName] = sourceValue
		if isGenerated {
			r.Fields[fieldNameCertainty] = certaintyGenerated
		} else {
			r.Fields[fieldNameCertainty] = certaintyPrefilled
		}
		return
	}

	// skip the update procedure if the certainty value is set to human or verified
	if currentCertainty == certaintyHuman || currentCertainty == certaintyVerified {
		return
	}

	// for all other reasons update the field
	r.IsUpdated = true
	r.Fields[fieldName] = sourceValue
	if isGenerated {
		r.Fields[fieldNameCertainty] = certaintyGenerated
	} else {
		r.Fields[fieldNameCertainty] = certaintyPrefilled
	}
	return
}

// handlerFunc is used to handle specific fields
type handlerFunc func(value string) (interface{}, bool)
