package reverse

func Reverse(s string) string {
	if s == "" {
		return ""
	}

	rune := []rune(s)
	reverse := ""

	for i := len(rune) - 1; i >= 0; i-- {
		reverse += string(rune[i])
	}

	return reverse
}
