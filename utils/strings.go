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
	rStr := []rune(str)
	rOld := []rune(old)
	rNew := []rune(new)

	i := 0
	for i < len(rStr) {
		if i+len(rOld) <= len(rStr) && string(rStr[i:i+len(rOld)]) == old {
			result += string(rNew)
			i += len(rOld)
		} else {
			result += string(rStr[i])
			i++
		}
	}

	return result
}

func Split(str, sep string) []string {
	result := []string{""}

	rStr := []rune(str)
	rSep := []rune(sep)

	i := 0
	for i < len(rStr) {
		if i+len(rSep) <= len(rStr) && string(rStr[i:i+len(rSep)]) == sep {
			result = append(result, "")
			i += len(rSep)
		} else {
			result[len(result)-1] += string(rStr[i])
			i++
		}
	}
	return result
}

func TrimSlice(slice []string) []string {
	return Split(Join(slice, " "), " ")
}
