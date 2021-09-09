package util

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)



var transformChain = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func transformString(s string) string {
	result, _, _ := transform.String(transformChain, s)
	return result
}

func GenerateSlug(title string) string {
	title = transformString(title)
	words := strings.Split(title, " ")
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return strings.Join(words, "-")
}
