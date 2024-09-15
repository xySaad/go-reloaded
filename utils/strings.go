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

	i := 0
	for i < len(str) {
		if i+len(old) <= len(str) && str[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(str[i])
			i++
		}
	}

	return result
}

func Split(str, sep string) []string {
	result := []string{""}

	i := 0
	for i < len(str) {
		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {
			result = append(result, "")
			i += len(sep)
		} else {
			result[len(result)-1] += string(str[i])
			i++
		}
	}
	return result
}
