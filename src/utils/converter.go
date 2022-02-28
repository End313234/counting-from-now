package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertTimestamp(seconds uint) string {
	attrs := [][]string{{"hour", "3600"}, {"minute", "60"}, {"second", "1"}}
	results := make([]string, 0)

	for _, value := range attrs {
		name, valueInSeconds := value[0], value[1]
		parsedValue, _ := strconv.Atoi(valueInSeconds)
		if int(seconds) >= parsedValue {
			var addS string

			result, rest := seconds/uint(parsedValue), seconds%uint(parsedValue)
			if result > 1 {
				addS = "s"
			}

			seconds = rest
			results = append(results, fmt.Sprintf("%s %s%s", fmt.Sprint(result), name, addS))
		}
	}

	finalAdition := fmt.Sprintf("and %s", results[len(results)-1])
	if len(results) == 1 {
		finalAdition = results[len(results)-1]
	}

	return fmt.Sprintf("%s %s", strings.Join(results[:len(results)-1], ","), finalAdition)
}
