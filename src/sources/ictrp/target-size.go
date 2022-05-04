package ictrp

import (
	"strconv"
	"strings"
)

func getNEnrollmentFn(print bool) func(string) (interface{}, bool) {
	// convert the given target size to an integer
	return func(value string) (interface{}, bool) {
		value = strings.ToLower(value)

		// if the string is an integer return that directly
		if nEnrollment, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
			return nEnrollment, true
		}

		if strings.Contains(value, ";") {
			// target_size of form: two groups of: 80;
			targets := strings.Split(value, ";")

			targetSize := 0
			for _, target := range targets {
				// if the group is not of structure: a group: 79;
				if !strings.Contains(target, ":") {
					continue
				}

				group := strings.Split(target, ":")

				groupName := group[0]
				groupSize, err := strconv.Atoi(strings.TrimSpace(group[1]))
				if err != nil {
					continue
				}

				multiplyer := 1
				if strings.HasPrefix(groupName, "two ") {
					multiplyer = 2
				}
				if strings.HasPrefix(groupName, "three ") {
					multiplyer = 3
				}
				if strings.HasPrefix(groupName, "four ") {
					multiplyer = 4
				}
				if strings.HasPrefix(groupName, "five ") {
					multiplyer = 5
				}

				targetSize += (multiplyer * groupSize)
			}

			// if no group found return nil
			if targetSize == 0 {
				return nil, true
			}

			return targetSize, true
		}

		if strings.Contains(value, "(") && strings.Contains(value, ")") {
			targets := strings.Split(value, "(")
			targetSizeString := strings.TrimSpace(targets[0])

			groupSize, err := strconv.Atoi(targetSizeString)
			if err == nil {
				return groupSize, true
			}
		}

		return nil, false
	}
}
