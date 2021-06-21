package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	replacer := strings.NewReplacer("-", " ", "_", "")
	replaced := replacer.Replace(s)
	slice := strings.Split(replaced, " ")
	abbreviate := ""

	for _, item := range slice {
		if len(item) > 0 {
			abbreviate += item[0:1]
		}
	}

	return strings.ToUpper(abbreviate)
}
