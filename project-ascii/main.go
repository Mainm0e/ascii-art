package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ascii-art/project-ascii/printArg"
	"github.com/ascii-art/project-ascii/Options"
)

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//fineTheme gona take string from os.Args[2] and check what file we gona use
func fineTheme(s string) string {

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
func fineOption(s string)int{
	option := s
	var index int

	switch option{
	case "--output":
		index = 1
	case "--align":
		index = 2
	case "--color":
		index =3
	default:
		index = 0
	}

	return index
}

func main() {

	if len(os.Args) == 1 {
		os.Exit(0)
	}

	if !isASCII(os.Args[1]) {
		fmt.Println("Non-ASCII characters")
		os.Exit(0)
	}
	//Looking at Args 2 if there have right Theme
	Theme := "ascii-art/standard.txt"
	if len(os.Args) > 2 {
		Theme = fineTheme(os.Args[2])
	}
	// Just Print too terminal but not for option
	if len(os.Args) < 5 {
		A1 := os.Args[1]
		sliceB := strings.Split(A1, "\\n")
		for i := 0; i < len(sliceB); i++ {
			printArg.Printer(sliceB[i], Theme)
		}
	}
	Options.outputArg()
	/*//_______________________________________________\\*\
	\*\\                                               //*/
	//Looking at Args 3 if there have right Option
	/*
	var object string
	if len(os.Args) == 4 {
		A3 := os.Args[3]
		sliceA := strings.Split(A3, "=")
		if len(sliceA) == 2 {
			object = sliceA[1]
			indexOption := fineOption(sliceA[0])
			switch indexOption{
			case 0:
				fmt.Println("Option incorrect")
			case 1:
				Options.outputArg(os.Args[1],Theme,object)
			case 2:
				Options.alignArg(os.Args[1],Theme,object)
			case 3:
				Options.colorArg(os.Args[1],Theme,object)
			default:
				fmt.Println("Option incorrect")
			} 
		}
	}
	*/

	//sliceA := strings.Split(os.Args[3], "=")



}
