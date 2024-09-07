package utils

import (
	"fmt"
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
		result = append(result[:i], result[i+1:]...)
		command := matches[1]
		index, _ := strconv.Atoi(matches[2])
		if index == 0 {
			index = 1
		}
		for j := 0; j > 0; j++ {
			if result[i-j] != " " && result[i-j] != "," {
				break
			}
			index++
		}
		switch command {
		case "hex":
			for rep := 1; rep <= index; rep++ {
				result[i-rep] = Hex(result[i-rep])
			}

		case "bin":
			for rep := 1; rep <= index; rep++ {
				result[i-rep] = Bin(result[i-rep])
			}
		case "cap":
			for rep := 1; rep <= index; rep++ {
				result[i-rep] = Cap(result[i-rep])
			}
		}
	}
	return strings.Join(result, " ")
}

func Hex(str string) string {
	data, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strconv.Itoa(int(data))
}

func Bin(str string) string {
	data, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strconv.Itoa(int(data))
}
func Cap(str string) string {
	println('A' - 'a')
	slice := []rune(str)
	slice[0] = slice[0] - 32
	return string(slice)
}
