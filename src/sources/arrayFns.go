package sources

import (
	"fmt"
	"strings"
)

// ToArray returns an database array out of an go array
func ToArray(arr []string) (interface{}, bool) {
	return fmt.Sprintf(`[%s]`, strings.Join(arr, ",")), true
}

// ToList returns a database list out of an go slice
func ToList(arr []string) (interface{}, bool) {
	return strings.Join(arr, ";"), true
}

// SetGenerated will return the same interface but returns that it is generated
func SetGenerated(m string) (interface{}, bool) {
	return m, true
}
