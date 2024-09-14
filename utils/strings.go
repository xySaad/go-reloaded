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
