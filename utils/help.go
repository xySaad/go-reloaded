package utils

import (
	"fmt"
	"os"
)

func PrintWarn(warn string) {
	switch warn {
	case "noArgs":
		fmt.Fprintln(os.Stderr, "[go reloaded] Please provide input and output files")
	case "noOutput":
		fmt.Fprintln(os.Stderr, "[go reloaded] Please provide output file")
	case "usage":
		fmt.Fprintln(os.Stderr, "[go reloaded] Usage: go run . input.txt output.txt")
	case "manyArgs":
		fmt.Fprintln(os.Stderr, "[go reloaded] Too many arguments")
	}
}
