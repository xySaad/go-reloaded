package main

import (
	"fmt"
	"go_reloaded/utils"
	"os"
	"strings"
)

func main() {
	args := os.Args

	// check usage
	if len(args) < 3 {
		if len(args) < 2 {
			utils.PrintWarn("noArgs")
		} else {
			utils.PrintWarn("noOutput")
		}
		utils.PrintWarn("usage")
		return
	}

	input := args[1]
	data, err := os.ReadFile(input)

	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}

	output := args[2]

	plainTxt := string(data)
	lines := strings.Split(utils.Replace(plainTxt, "\r\n", "\n"), "\n")
	convertedTxt := ""

	for i, line := range lines {
		formatedTxt := utils.FormatTxt(line)
		convertedTxt += utils.Convert(formatedTxt)
		if i < len(lines)-1 {
			convertedTxt += "\n"
		}
	}

	outputDir := "./output"

	mkdirErr := os.MkdirAll(outputDir, 0755)

	if mkdirErr != nil {
		fmt.Fprintln(os.Stderr, "Error creating output directory:", err)
		return
	}

	writeErr := os.WriteFile(outputDir+"/"+output, []byte(convertedTxt), 0644)

	if writeErr != nil {
		fmt.Fprintln(os.Stderr, "Error writing file:", writeErr)
	} else {
		fmt.Println("File", output, "has been written successfully")
	}
}
