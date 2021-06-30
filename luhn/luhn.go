package luhn

import (
	"strconv"
	"strings"
)

func Valid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	trim := strings.ReplaceAll(s, " ", "")
	sum := 0
	doubleIndex := len(trim) - 2

	for i := len(trim) - 1; i > 0; i-- {
		b := trim[i]
		toInt, err := strconv.Atoi(string(b))
		if err != nil {
			return false
		}

		if i == doubleIndex {
			doubleIndex = i - 2
			toInt = toInt * 2
			if toInt > 9 {
				toInt = toInt - 9
			}
		}
		sum += toInt
	}
	return sum%10 == 0
}
