package utils

func Join(slice []string, sep string) string {
	output := ""
	start := false
	spaceNum := 0

	for _, v := range slice {
		if !start && v == "" {
			spaceNum++
			if spaceNum%4 == 0 {
				output += "\t"
			}
		}

		if v != "" {
			output += v + sep
			start = true
		}
	}

	if len(output)-1 >= 0 {
		output = output[:len(output)-1]
	}

	return output
}

func Replace(str, old, new string) string {
	result := ""
	oldText := []rune(old)
	text := []rune(str)

	for i, char := range text {
		pass := true

		if i+len(oldText) < len(text) {
			for j, oldChar := range oldText {
				if text[i+j] != oldChar {
					pass = false
				}
			}
		}

		if pass {
			result += new
		} else {
			result += string(char)
		}

	}

	return result
}
