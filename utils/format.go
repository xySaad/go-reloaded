package utils

import (
	"regexp"
)

func FormatTxt(txt string) []string {
	result := txt

	re := regexp.MustCompile(` \((hex|bin|up|low|cap), (\d+)\)[,|\.\.\.|,|!|\?|:]*`)
	result = re.ReplaceAllString(result, `$3 ($1-$2)`)

	punctuationRe := regexp.MustCompile(`[ ]*([,.!?:]+)[ ]*`)
	result = punctuationRe.ReplaceAllString(result, `$1`)

	punctRe2 := regexp.MustCompile(`([,\.!\?:]+)`)
	result = punctRe2.ReplaceAllString(result, `$1 `)

	apstropheRe := regexp.MustCompile(`(\w)'(\w)`)
	result = apstropheRe.ReplaceAllString(result, `$1’$2`)

	quotationRe := regexp.MustCompile(`'[ ]*(.*?)[ ]*'`)
	result = quotationRe.ReplaceAllString(result, `'$1'`)

	apstropheRe = regexp.MustCompile(`(\w)’(\w)`)
	result = apstropheRe.ReplaceAllString(result, `$1'$2`)

	return Split(result, " ")
}

func A2An(str string) string {
	a2anRe := regexp.MustCompile(`\b(a|A)[ ]+([aeiouhAEIOUH])`)

	for {
		newStr := a2anRe.ReplaceAllString(str, "${1}n $2")
		if newStr == str {
			break
		}
		str = newStr
	}

	return str
}
