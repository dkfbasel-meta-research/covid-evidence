package ictrp

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

	// translate months from german to english
	value = strings.ReplaceAll(value, "Mär", "Mar")
	value = strings.ReplaceAll(value, "Mai", "May")
	value = strings.ReplaceAll(value, "Okt", "Oct")
	value = strings.ReplaceAll(value, "Dez", "Dec")

	asTime, err := time.Parse("January 2, 2006", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("2 January 2006", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("January 2006", value)
	if err == nil {
		dt := asTime.Format("2006-01")
		return fmt.Sprintf("%s-15", dt), true
	}

	asTime, err = time.Parse("2-Jan-06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("Jan 2, 06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("02-Jan-06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("2. Jan 06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("06-Jan", value)
	if err == nil {
		dt := asTime.Format("2006-01")
		return fmt.Sprintf("%s-15", dt), true
	}

	asTime, err = time.Parse("02.01.06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("01/02/06", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("02/01/2006", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("2006/01/02", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("01-02-2006", value)
	if err == nil {
		return asTime.Format("2006-01-02"), true
	}

	asTime, err = time.Parse("2006-02", value)
	if err == nil {
		dt := asTime.Format("2006-01")
		return fmt.Sprintf("%s-15", dt), true
	}

	if strings.ToLower(value) == "na" {
		return "", true
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
