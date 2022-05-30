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
	variables := map[string]string{}

	for _, line := range fileLines {
		// Checking if the line starts with `//`
		if strings.HasPrefix(line, "//") {
			continue
		} else if strings.Contains(line, "=") {
			varName := strings.Split(line, "=")[0]
			varValue := strings.Split(line, "=")[1]
			if strings.HasPrefix(varValue, "\"") && strings.HasSuffix(varValue, "\"") {
				variables[string(varName)] = string(varValue[1 : len(varValue)-1])
			} else {
				variables[string(varName)] = string(varValue)
			}
		} else if strings.HasPrefix(line, "print") {
			args := (strings.Split(line, " "))[1]
			// Checking if the string has quotes at the beginning and end.
			if strings.HasPrefix(args, "\"") && strings.HasSuffix(args, "\"") {
				fmt.Println(args[1 : len(args)-1])
				// It's checking if the variable exists in the map.
			} else {
				// every key in the map is a string.
				for key, value := range variables {
					// first we check if the variable exists in the map.
					if strings.Contains(args, key) {
						// if it exists we print out the value of the variable.
						fmt.Println(variables[value])
					}
				}
			}
		}
	}

	os.WriteFile("variables.txt", []byte(fmt.Sprintf("%v", variables)), 0644)
}
