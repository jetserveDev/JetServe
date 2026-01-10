package main

import "github.com/fatih/color"

func instruction() {

	var content string = `
                     INSTRUCTION
╔═════════════════════════════════════════════════════╗
║      start server - to start creating server        ║
║                                                     ║
║      exit - close programm(Data will not be saved)  ║
║                                                     ║
║      cancel - to cancel an action                   ║
║                                                     ║
║     For detailed instructions, please visit the     ║
║                website: http://example.com          ║
╚═════════════════════════════════════════════════════╝

All commands for managing the file system are the same(ls,cd,pwd)
     
`
	color.Cyan(content)

}
