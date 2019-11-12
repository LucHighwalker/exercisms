// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

/*
Bob responds with various snarky remarks
silence: Fine. Be that way!
yelling: Whoa, chill out!
question: Sure.
yell question: Calm down, I know what I'm doing!
default: Whatever.
*/
func Hey(s string) string {
	str := strings.TrimSpace(s)

	if len(str) < 1 {
		return "Fine. Be that way!"
	}

	yell := yelling(str)
	quest := question(str)

	if yell && quest {
		return "Calm down, I know what I'm doing!"
	} else if yell {
		return "Whoa, chill out!"
	} else if quest {
		return "Sure."
	}

	return "Whatever."
}

func yelling(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}

func question(s string) bool {
	return s[len(s)-1:] == "?"
}
