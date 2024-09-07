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

	if strings.HasSuffix(output, ".go") {
		print("Can't process a go file, please enter a text file path")
		return
	}

	plainTxt := string(data)

	formatedTxt := utils.FormatTxt(plainTxt)
	convertedTxt := utils.Convert(formatedTxt)
	writeErr := os.WriteFile(output, []byte(convertedTxt), 0644)
	if writeErr != nil {
		fmt.Println("Error writing file:", writeErr)
	} else {
		fmt.Println("File", output, "has been written successfully")
	}
}
