package clock

import "fmt"

type clock struct {
	hours   int
	minutes int
}

func New(hours, minutes int) clock {
	fH, fM := format(hours, minutes)

	return clock{
		hours:   fH,
		minutes: fM,
	}
}

func (c clock) String() string {
	hoursToString := fmt.Sprintf("%d", c.hours)
	if len(hoursToString) == 1 {
		hoursToString = "0" + hoursToString
	}

	minutesToString := fmt.Sprintf("%d", c.minutes)
	if len(minutesToString) == 1 {
		minutesToString = "0" + minutesToString
	}

	return hoursToString + ":" + minutesToString
}

func (c clock) Add(minutes int) clock {
	fH, fM := format(c.hours, c.minutes+minutes)

	return clock{
		hours:   fH,
		minutes: fM,
	}
}

func (c clock) Subtract(minutes int) clock {
	fH, fM := format(c.hours, c.minutes-minutes)

	return clock{
		hours:   fH,
		minutes: fM,
	}
}

func format(hours, minutes int) (int, int) {
	hoursFormated := hours
	minutesFormated := minutes

	if minutes > 59 {
		minutesFormated = minutes % 60
		hoursFormated += int(minutes / 60)
	}
	if minutes < 0 && minutes >= -60 {
		minutesFormated = minutes + 60
		hoursFormated--
	}
	if minutes < -60 {
		minutesFormated = minutes % 60 * -1
		hoursFormated += int(minutes/60) * -1
	}

	if hoursFormated >= 24 {
		hoursFormated = hoursFormated % 24
	}
	if hoursFormated < 0 && hoursFormated > -24 {
		hoursFormated = hoursFormated + 24
	}
	if hoursFormated <= -24 {
		hoursFormated = hoursFormated % 24 * -1
	}

	return hoursFormated, minutesFormated
}
