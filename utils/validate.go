package utils

import (
	"strings"
	"unicode"
)

func IsValidMarkdown(filename string) bool {
	if !(strings.HasSuffix(filename, ".md") || strings.HasSuffix(filename, ".markdown")) {
		return false
	}

	name := filename[:strings.LastIndex(filename, ".")]
	for _, ch := range name {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != ' ' && ch != '-' && ch != '_' {
			return false
		}
	}

	return true
}
