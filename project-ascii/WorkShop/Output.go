package WorkShop

import (
	"bufio"
	"log"
	"os"
)

func Output(s, t, o string) {
	String := s
	Theme := t
	FileName := o
	var NewString string

	file, err := os.Open(Theme)

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
	/*
		splitString := strings.Split(os.Args[1], "\\n")

		var asciiMap = make(map[int][]string)

		for i := 0 ; i < len(splitString); i++ {

		}

		/*
			for i := 0; i < len(String); i++ {
				letter := String[i]


				for numb, eachline := range txtlines {

				}
			}
	*/
	NewString = "hi"
	output := []byte(NewString)
	MakeFile := os.WriteFile(FileName, output, 0644)
	check(MakeFile)
}
