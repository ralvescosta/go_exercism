package clock

type clock struct {
	hours   int
	minutes int
}

func New(hours, minutes int) clock {
	if minutes > 60 {
		minutes = minutes % 60
		hours += int(minutes / 60)
	}
	if minutes < 0 && minutes > -60 {
		minutes = minutes + 60
		hours--
	}
	if minutes <= -60 {
		minutes = minutes % 60 * -1
		hours += int(minutes/60) * -1
	}

	if hours >= 24 {
		hours = hours % 24
	}
	if hours < 0 && hours > -24 {
		hours = hours + 24
	}
	if hours <= -24 {
		hours = hours % 24 * -1
	}

	return clock{
		hours:   hours,
		minutes: minutes,
	}
}

func (clock) String() string {
	return ""
}

func (clock) Add(minutes int) clock {
	return clock{}
}

func (clock) Subtract(minutes int) clock {
	return clock{}
}
