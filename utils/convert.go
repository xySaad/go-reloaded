package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func Convert(txt []string) string {
	return convertHelper(txt, 0)
}

func convertHelper(txt []string, index int) string {
	if index >= len(txt) {
		return Join(txt, " ")
	}

	v := txt[index]

	// re := regexp.MustCompile(`\((hex|bin|up|low|cap)-?(\d+)?\)$`)
	quantifier := 0
	modifier := ""

	if v == "(hex)" || v == "(bin)" || v == "(low)" || v == "(cap)" || v == "(up)"{
		modifier = v[1:len(v)-1]
		quantifier = 1
		// Remove the command from the slice
		txt = append(txt[:index], txt[index+1:]...)
	} else if (v == "(low," || v == "(up," || v == "(cap,") && txt[index+1][len(txt[index+1])-1] == ')' {
		num, isValidNumber := Atoi(txt[index+1][:len(txt[index+1])-1])
		if isValidNumber && num >= 0 {
			modifier = v[1:len(v)-1]
			quantifier = num
			// Remove the command from the slice
			txt = append(txt[:index], txt[index+2:]...)
		}
	}

	if modifier == "" {
		return convertHelper(txt, index+1)
	}

	// Apply the modifier to previous elements
	for rep := 1; rep <= quantifier; rep++ {
		if index-rep >= 0 {
			switch modifier {
			case "hex":
				txt[index-rep] = Hex(txt[index-rep])
			case "bin":
				txt[index-rep] = Bin(txt[index-rep])
			case "cap":
				txt[index-rep] = Cap(txt[index-rep])
			case "up":
				txt[index-rep] = strings.ToUpper(txt[index-rep])
			case "low":
				txt[index-rep] = strings.ToLower(txt[index-rep])
			}
		} else {
			break
		}
	}

	// Continue processing with the modified text
	return convertHelper(txt, index)
}

func Hex(str string) string {
	data, err := strconv.ParseInt(str, 16, 0)
	if err != nil {
		// fmt.Println(str, "-> Error:", err)
		return str
	}
	return strconv.Itoa(int(data))
}

func Bin(str string) string {
	data, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		// fmt.Println(str, "-> Error:", err)
		return str
	}
	return strconv.Itoa(int(data))
}
func Cap(str string) string {
	slice := []rune(str)

	if len(slice) <= 0 {
		return str
	}
	start := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] == '\'' || slice[i] == '(' {
			start = 0
		} else {
			start = i
			break
		}
	}

	slice[start] = unicode.ToUpper(slice[start])

	for i := start + 1; i < len(slice); i++ {
		slice[i] = unicode.ToLower(slice[i])
	}
	return string(slice)
}

func Atoi(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	return num, err == nil
}
