// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func verifyString(s string) (bool, bool, bool) {
	hasNumber := false
	hasLetter := false
	hasDigit := false

	for _, r := range s {
		if unicode.IsNumber(r) {
			hasNumber = true
		}
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasNumber, hasLetter, hasDigit
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	hasQuestionMark := remark[len(remark)-1] == '?'
	isUpperCase := strings.ToUpper(remark) == remark
	hasNumber, hasLetter, hasDigit := verifyString(remark)
	isNumberSequency := hasNumber && !hasLetter && hasDigit

	if (isUpperCase && !isNumberSequency) && !hasQuestionMark {
		return "Whoa, chill out!"
	}

	if (isUpperCase && !isNumberSequency) && hasQuestionMark {
		return "Calm down, I know what I'm doing!"
	}

	if hasQuestionMark {
		return "Sure."
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
