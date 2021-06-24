package raindrops

import "fmt"

func Convert(rain int) string {
	drops := ""

	if rain%3 == 0 {
		drops += "Pling"
	}
	if rain%5 == 0 {
		drops += "Plang"
	}
	if rain%7 == 0 {
		drops += "Plong"
	}
	if len(drops) == 0 {
		return fmt.Sprintf("%d", rain)
	}

	return drops
}
