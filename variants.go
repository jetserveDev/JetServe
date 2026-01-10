package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Variants() {
	var variantsOfServers = `
╔════════════════════════════════════════════════╗
║                 SERVER MODE:                   ║
║                                                ║
║    local - LOCALHOST (only this computer)      ║
║                                                ║
║    public - PUBLIC (accessible from internet)  ║
║                                                ║
╚════════════════════════════════════════════════╝
`
	color.Cyan(variantsOfServers)
	reader := bufio.NewReader(os.Stdin)
	running := true

	for running {
		fmt.Print("CREATING SERVER>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
		}
		input = strings.TrimSpace(input)

		if input == "cancel" {
			color.Red("CANCELED")
			running = false
		} else if input == "local" {
			localServer()
		} else if input == "public" {
			public()
		}

	}
}
