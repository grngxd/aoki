package main

import (
	run "aoki/run"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		if args[0] == "run" {
			if len(args[1]) != 0 && strings.HasSuffix(args[1], ".aoki") {
				run.Run(args[1])
			} else {
				fmt.Println("Please specify a .aoki file.")
			}
		} else {
			fmt.Printf("The command \"%s\" does not exist.", args[0])
		}
	} else {
		fmt.Println("Aoki")
	}
}
