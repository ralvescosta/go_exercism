package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	result := FreqMap{}
	c := make(chan FreqMap)

	for _, text := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(text)
	}

	for range texts {
		for i, m := range <-c {
			result[i] += m
		}
	}

	return result
}
