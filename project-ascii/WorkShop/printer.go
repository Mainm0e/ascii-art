package WorkShop

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Printer(s1, s2 string) {

	argTheme := s2

	file, err := os.Open(argTheme)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	var l1, l2, l3, l4, l5, l6, l7, l8 string
	for i := 0; i < len(s1); i++ {

		var y int
		letter := s1[i]
		DecN := int(letter - 32)

		for i := 0; i < DecN; i++ {
			y = y + 9
		}

		stringNumb := y + 1

		for i := 1; i <= 8; i++ {
			line := i
			for numb, eachline := range txtlines {
				if numb == stringNumb {
					switch line {
					case 1:
						l1 = l1 + eachline
					case 2:
						l2 = l2 + eachline
					case 3:
						l3 = l3 + eachline
					case 4:
						l4 = l4 + eachline
					case 5:
						l5 = l5 + eachline
					case 6:
						l6 = l6 + eachline
					case 7:
						l7 = l7 + eachline
					case 8:
						l8 = l8 + eachline
					default:
					}
				}

			}
			stringNumb++
		}
	}
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
	fmt.Println(l5)
	fmt.Println(l6)
	fmt.Println(l7)
	fmt.Println(l8)

}
