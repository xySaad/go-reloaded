package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func Convert(txt []string) string {
	return convertHelper(txt, 0)
}

func convertHelper(txt []string, index int) string {
	if index >= len(txt) {
		return Join(txt, " ")
	}

	v := txt[index]
	re := regexp.MustCompile(`\((hex|bin|up|low|cap)-?(\d+)?\)`)
	matches := re.FindStringSubmatch(v)

	if len(matches) < 1 {
		return convertHelper(txt, index+1)
	}

	modifier := matches[1]
	quantifier, _ := strconv.Atoi(matches[2])
	if quantifier == 0 {
		quantifier = 1
	}

	// Remove the command from the slice
	txt = append(txt[:index], txt[index+1:]...)

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
		if slice[i] != '\'' {
			start = i
			break
		}
	}
	if slice[start] <= 'z' && slice[start] >= 'a' {
		slice[start] = slice[start] - 32
	}
	for i := start + 1; i < len(slice); i++ {
		if slice[i] <= 'Z' && slice[i] >= 'A' {
			slice[i] = slice[i] + 32
		}
	}
	return string(slice)
}
