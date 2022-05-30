package aoki

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(filePath string) {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		if strings.HasPrefix(line, "//") {
			continue
		} else if strings.HasPrefix(line, "print") {
			args := (strings.Split(line, " "))[1]
			if strings.HasPrefix(args, "\"") && strings.HasSuffix(args, "\"") {
				fmt.Println(args[1 : len(args)-1])
			} else {
				fmt.Println(args)
			}
		}
	}
}
