package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func Convert(txt []string) string {

	result := txt

	for i, v := range txt {
		re := regexp.MustCompile(`\((hex|bin|up|low|cap)\-?(\d*)?\)`)
		matches := re.FindStringSubmatch(v)
		if len(matches) < 1 {
			continue
		}
		if len(result) == i-1 {
			result = result[:i-1]
			break
		} else {
			result = append(result[:i], result[i+1:]...)
		}
		command := matches[1]
		index, _ := strconv.Atoi(matches[2])
		if index == 0 {
			index = 1
		}

		switch command {
		case "hex":
			for rep := 1; rep <= index; rep++ {
				if i-rep >= 0 {
					result[i-rep] = Hex(result[i-rep])
				}
			}

		case "bin":
			for rep := 1; rep <= index; rep++ {
				if i-rep >= 0 {
					result[i-rep] = Bin(result[i-rep])
				}
			}
		case "cap":
			for rep := 1; rep <= index; rep++ {
				if i-rep >= 0 {
					result[i-rep] = Cap(result[i-rep])
				}
			}
		case "up":
			for rep := 1; rep <= index; rep++ {
				if i-rep >= 0 {
					result[i-rep] = strings.ToUpper(result[i-rep])
				}
			}
		case "low":
			for rep := 1; rep <= index; rep++ {
				if i-rep >= 0 {
					result[i-rep] = strings.ToLower(result[i-rep])
				}
			}
		}
	}
	return strings.Join(result, " ")
}

func Hex(str string) string {
	data, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		// fmt.Println(str, "-> Error:", err)
		return str
	}
	return strconv.Itoa(int(data))
}

func Bin(str string) string {
	data, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		// fmt.Println(str, "-> Error:", err)
		return str
	}
	return strconv.Itoa(int(data))
}
func Cap(str string) string {
	slice := []rune(str)
	if slice[0] <= 'z' && slice[0] >= 'a' {
		slice[0] = slice[0] - 32

	}
	return string(slice)
}
