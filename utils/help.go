package utils

func PrintWarn(warn string) {
	switch warn {
	case "noArgs":
		print("[go reloaded] Please provide input and output files\n")
	case "noOutput":
		print("[go reloaded] Please provide output file\n")
	case "usage":
		print("[go reloaded] Usage: go run . input.txt output.txt")
	}
}
