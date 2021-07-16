package twelve

import "strings"

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i)

		if i != 12 {
			song += "\n"
		}
	}
	return song
}

func Verse(day int) string {
	verse := strings.ReplaceAll("On the $d day of Christmas my true love gave to me: ", "$d", days[day-1])

	if day == 1 {
		return verse + dayVerses[0] + "."
	}

	for i := day - 1; i >= 0; i-- {
		if i == 0 && day >= 2 {
			verse += "and " + dayVerses[i] + "."
		}
		if i > 0 && day > 0 {
			verse += dayVerses[i] + ", "
		}
	}
	return verse
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var dayVerses = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}
