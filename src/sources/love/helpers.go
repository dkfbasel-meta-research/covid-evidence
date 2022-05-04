package love

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"dkfbasel.ch/covid-evidence/logger"
)

// isEqual will determine if to entries are equal
func isEqual(value1, value2 string) (fullEqual, partialEqual bool) {

	if value1 == value2 {
		fullEqual = true
	}

	if clean(value1) == clean(value2) {
		partialEqual = true
	}

	return fullEqual, partialEqual

}

func clean(value string) string {

	s := []byte(value)

	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			s[j] = b
			j++
		}
	}
	return strings.ToLower(string(s[:j]))
}

func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	return asString(value) == ""
}

// return string representation of value
func asString(value interface{}) string {

	if value == nil {
		return ""
	}

	switch t := value.(type) {
	case int, int32, int64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%.0f", t)
	case bool:
		return fmt.Sprintf("%t", t)
	default:
		return fmt.Sprintf("%s", t)
	}

}

// convert the given value to an iso date
func toIsoDate(value string) (interface{}, bool) {

	// 3/30/21 13:05
	asTime, err := time.Parse("1/2/06 15:04", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 26.04.2021 18:49
	asTime, err = time.Parse("02.01.2006 15:04", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 3/30/21 13:05
	asTime, err = time.Parse("1/02/06 15:04", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 10/12/20 12:28
	asTime, err = time.Parse("01/02/06 15:04", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 08/18/21 09:52 AM
	asTime, err = time.Parse("01/02/06 03:04 PM", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 2020-09-28T13:42:18Z
	asTime, err = time.Parse("2006-01-02T15:04:05Z", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 20-09-28T13:42:15Z
	asTime, err = time.Parse("06-01-02T15:04:05Z", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	// 9/20 12:01
	asTime, err = time.Parse("1/06 15:04", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	if err != nil {
		logger.Debug("could not parse date", logger.String("value", value))
	}

	return value, false
}

// toLowerCase will convert the value to lowercase
func toLowerCase(value string) (interface{}, bool) {
	return strings.ToLower(value), false
}

func toInt(value string) (interface{}, bool) {
	asInt, _ := strconv.Atoi(value)
	return asInt, false
}
