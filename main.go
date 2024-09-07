package main

import (
	"fmt"
	"go_reloaded/utils"
	"os"
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

	plainTxt := string(data)

	formatedTxt := utils.FormatTxt(plainTxt)
	convertedTxt := utils.Convert(formatedTxt)
	output := args[2]
	writeErr := os.WriteFile(output, []byte(convertedTxt), 0644)
	if writeErr != nil {
		fmt.Println("Error writing file:", writeErr)
	} else {
		fmt.Println("File", output, "has been written successfully")
	}
}
