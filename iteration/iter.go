package iteration

import (
	"strings"
)

func Repeat(text string, iterations int) string {
	var res strings.Builder
	for i := 0; i < iterations; i++ {
		res.WriteString(text)
	}

	return res.String()
}

func OwnClone(input string) string {
	clone := strings.Clone(input)
	return clone
}
