package main

import (
	run "aoki/run"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		if args[0] == "run" {
			run.Run(args[1])
		} else {
			fmt.Println("You are trying to run a command")
		}
	} else {
		fmt.Println("Aoki")
	}
}
