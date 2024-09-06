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
	plainTxt, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error Reading file:", err)
		return
	}

	// Debuging
	print(plainTxt)
}
