package checker

import "fmt"

func Ascii(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

//fineTheme gona take string from os.Args[2] and check what file we gona use
func FineTheme(s string) string {

	if s == "standard" {
		return "Ascii-art/standard.txt"
	} else if s == "shadow" {
		return "Ascii-art/shadow.txt"
	} else if s == "thinkertoy" {
		return "Ascii-art/thinkertoy.txt"
	} else {
		return "Ascii-art/standard.txt"
	}

}

// fineOption take option and look what option file need to go
func FineOption(s string) int {
	option := s
	var index int

	switch option {
	case "--output":
		index = 1
	case "--align":
		index = 2
	case "--color":
		index = 3
	default:
		index = 0
	}

	return index
}

func UsageErr() {
	fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
	fmt.Println()
	fmt.Println("EX: go run . something standard --output=<fileName.txt>")
}
