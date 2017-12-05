package main

import (
	"strings"
	"regexp"
	"fmt"
)

func diff(a, b string, unified bool) string {
	var ap, bp rune
	if unified {
		ap = '-'
		bp = '+'
	} else {
		ap = '<'
		bp = '>'
	}

	minus_a := fmt.Sprintln(regexp.MustCompile(`(^|\n)`).ReplaceAllString(strings.TrimRight(a, "\n"), fmt.Sprintf("$1%c ", ap)))
	plus_b := fmt.Sprintln(regexp.MustCompile(`(^|\n)`).ReplaceAllString(strings.TrimRight(b, "\n"), fmt.Sprintf("$1%c ", bp)))
	return fmt.Sprintf("%s%s", minus_a, plus_b)
}
