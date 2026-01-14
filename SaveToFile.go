package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func saveToFile() bool {
	color.Green("Would you like to save logs to file(yes/no)")
	running := true
	reader := bufio.NewReader(os.Stdin)
	var response bool
	for running {
		input, err := reader.ReadString('\n')
		if err != nil {
			color.Red("ERROR:", err)
			break
		}

		input = strings.TrimSpace(input)

		if input == "yes" || input == "y" {
			response = true
			running = false
		} else if input == "no" || input == "no" {
			response = false
			running = false
		}
	}
	return response
}
