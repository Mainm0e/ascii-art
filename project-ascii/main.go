package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/WorkShop"
	"github.com/checker"
)

func main() {

	if len(os.Args) != 4 {
		checker.UsageErr()
	} else {

		if !checker.Ascii(os.Args[1]) {
			fmt.Println("Non-ASCII characters")
			os.Exit(0)
		}
		//Looking at Args 2 if there have right Theme
		Theme := "ascii-art/standard.txt"
		if len(os.Args) > 2 {
			Theme = checker.FineTheme(os.Args[2])
		}
		// Just Print too terminal but not for option
		/*if len(os.Args) < 4 {
			A1 := os.Args[1]
			sliceB := strings.Split(A1, "\\n")
			for i := 0; i < len(sliceB); i++ {
				WorkShop.Printer(sliceB[i], Theme)
			}
		}*/
		//Looking at Args 3 if there have right Option
		var object string
		if len(os.Args) == 4 {
			A3 := os.Args[3]
			sliceA := strings.Split(A3, "=")
			if len(sliceA) == 2 {
				object = sliceA[1]
				indexOption := checker.FineOption(sliceA[0])
				switch indexOption {
				case 0:
					fmt.Println("Option incorrect")
				case 1:
					WorkShop.Output(os.Args[1], Theme, object)
				case 2:
					WorkShop.Align(os.Args[1], Theme, object)
				case 3:
					WorkShop.Color(os.Args[1], Theme, object)
				default:
					fmt.Println("Option incorrect")
				}
			}
		}
	}
}

//sliceA := strings.Split(os.Args[3], "=")
