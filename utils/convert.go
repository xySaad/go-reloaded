package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func Convert(txt []string) string {

	result := txt

	for i := 0; i < len(result); i++ {
		v := result[i]
		re := regexp.MustCompile(`\((hex|bin|up|low|cap)\-?(\d+)?\)`)
		matches := re.FindStringSubmatch(v)
		if len(matches) < 1 {
			continue
		}
		if len(result) <= i {
			result = result[:i-1]
			break
		} else {
			result = append(result[:i], result[i+1:]...)
		}
		modifier := matches[1]
		quantifier, _ := strconv.Atoi(matches[2])

		if quantifier == 0 {
			quantifier = 1
		}

		// fmt.Println("command:", modifier, "index:", quantifier)

		switch modifier {
		case "hex":
			for rep := 1; rep <= quantifier; rep++ {
				if i-rep >= 0 {
					// fmt.Println(modifier+"ing", result[i-rep])
					result[i-rep] = Hex(result[i-rep])
					i--
				} else {
					break
				}
			}

		case "bin":
			for rep := 1; rep <= quantifier; rep++ {
				if i-rep >= 0 {
					// fmt.Println(modifier+"ing", result[i-rep])
					result[i-rep] = Bin(result[i-rep])
					i--
				} else {
					break
				}
			}
		case "cap":
			for rep := 1; rep <= quantifier; rep++ {
				if i-rep >= 0 {
					// fmt.Println(modifier+"ing", result[i-rep])
					result[i-rep] = Cap(result[i-rep])
					i--
				} else {
					break
				}
			}
		case "up":
			for rep := 1; rep <= quantifier; rep++ {
				if i-rep >= 0 {
					// fmt.Println(modifier+"ing", result[i-rep])
					result[i-rep] = strings.ToUpper(result[i-rep])
					i--
				} else {
					break
				}
			}
		case "low":
			for rep := 1; rep <= quantifier; rep++ {
				if i-rep >= 0 {
					// fmt.Println(modifier+"ing", result[i-rep])
					result[i-rep] = strings.ToLower(result[i-rep])
					i--
				} else {
					break
				}
			}
		}
	}

	return Join(result, " ")
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

	if slice[0] <= 'z' && slice[0] >= 'a' {
		slice[0] = slice[0] - 32
	}
	for i := 1; i < len(slice); i++ {
		if slice[i] <= 'Z' && slice[i] >= 'A' {
			slice[i] = slice[i] + 32
		}
	}
	return string(slice)
}
