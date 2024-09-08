package utils

import (
	"regexp"
	"strings"
)

func FormatTxt(txt string) []string {
	result := txt

	commaRe := regexp.MustCompile(`[\s]*([,|\.\.\.|,|!|\?|:]+)[.]*[\s]*([,|\.\.\.|,|!|\?|:]*)`)
	result = commaRe.ReplaceAllString(result, `$1$2 `)

	quotationRe := regexp.MustCompile(`([^n])'[\s]*(.*?)[\s]*'`)
	result = quotationRe.ReplaceAllString(result, `$1'$2'`)

	a2anRe := regexp.MustCompile(`(a|A)([\s]+[a|e|i|o|u|h|A|E|I|O|U|H])`)
	result = a2anRe.ReplaceAllString(result, `${1}n $2`)
	re := regexp.MustCompile(` \((hex|bin|up|low|cap), (\d*)?\)([,|\.\.\.|,|!|\?|:]*)?`)
	result = re.ReplaceAllString(result, `$3 ($1-$2)`)

	return strings.Fields(result)
}
