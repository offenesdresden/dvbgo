package dvb

import (
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func normalize(input *string) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	*input, _, _ = transform.String(t, *input)
}
