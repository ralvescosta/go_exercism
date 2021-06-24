package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("wrong length")
	}

	hammingDistance := 0

	for i, c := range a {
		if c != rune(b[i]) {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
