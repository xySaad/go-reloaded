package utils

import (
	"regexp"
	"strings"
)

func FormatTxt(txt string) []string {
	result := txt

	commaRe := regexp.MustCompile(`[ ]*([,\.!\?:]+)[ ]*`)
	result = commaRe.ReplaceAllString(result, `$1`)

	apstropheRe := regexp.MustCompile(`(\w)'(\w)`)
	result = apstropheRe.ReplaceAllString(result, `$1’$2`)

	quotationRe := regexp.MustCompile(`'[ ]*(.*?)[ ]*'`)
	result = quotationRe.ReplaceAllString(result, `'$1'`)

	a2anRe := regexp.MustCompile(`(^a|^A| a| A)([ ]+[a|e|i|o|u|h|A|E|I|O|U|H])`)
	result = a2anRe.ReplaceAllString(result, `${1}n $2`)

	re := regexp.MustCompile(` \((hex|bin|up|low|cap)(?:, (\d+)?)?\)([,|\.\.\.|,|!|\?|:]*)?`)
	result = re.ReplaceAllString(result, `$3 ($1-$2)`)

	return strings.Split(result, " ")
}
