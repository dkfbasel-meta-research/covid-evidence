package stores

import "fmt"

// AsString will return a string representation of the given value
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
