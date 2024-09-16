package utils

import (
	"regexp"
)

func FormatTxt(txt string) []string {
	result := txt

	re := regexp.MustCompile(` \((hex|bin|up|low|cap)(?:, (\d+)?)?\)([,|\.\.\.|,|!|\?|:]*)?`)
	result = re.ReplaceAllString(result, `$3 ($1-$2)`)

	punctuationRe := regexp.MustCompile(`[ ]*([,\.!\?:]+)[ ]*`)
	result = punctuationRe.ReplaceAllString(result, `$1`)

	punctRe2 := regexp.MustCompile(`([,\.!\?:]+)`)
	result = punctRe2.ReplaceAllString(result, `$1 `)

	apstropheRe := regexp.MustCompile(`(\w)'(\w)`)
	result = apstropheRe.ReplaceAllString(result, `$1’$2`)

	quotationRe := regexp.MustCompile(`'[ ]*(.*?)[ ]*'`)
	result = quotationRe.ReplaceAllString(result, `'$1'`)

	a2anRe := regexp.MustCompile(`(^a|^A| a| A)([ ]+[a|e|i|o|u|h|A|E|I|O|U|H])`)
	result = a2anRe.ReplaceAllString(result, `${1}n$2`)

	return Split(result, " ")
}
