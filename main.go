package main

import (
	"fmt"

	"github.com/fatih/color"

	"bufio"

	"os"

	"strings"
)

func main() {
	var skull = `
                            ,--.
                           {    }
                           K,   }
                          /  ~Y'
                     ,   /   /
                    {_'-K.__/
                      '/-.__L._
                      /  ' /'\_}
                     /  ' /
             ____   /  ' /
      ,-'~~~~    ~~/  ' /_
    ,'             '''~  ',		     
   (                        Y
  {                         I            ____.       __          		 
 {      -                    ',         |    | _____/  |_ 		 
 |       ',                   )         |    |/ __ \   __\		 
 |        |   ,..__      __. Y      /\__|    \  ___/|  | 	
 |    .,_./  Y ' / ^Y   J   )|      \________|\___  >__|  
 \           |' /   |   |   ||                    \/ 
  \          L_/    . _ (_,.'(       _________        		
   \,   ,      ^^""' / |      )     /   _____/ ______________  __ ____  		 
     \_  \          /,L]     /      \_____  \_/ __ \_  __ \  \/ // __ \ 		
       '-_~-,       ' '   ./        /        \  ___/|  | \/\   /\  ___/ 		       
          ''{_            )        /_______  /\___  >__|    \_/  \___  >
              ^^\..___,.--'                \/     \/                 \/ 
`

	color.Green(skull)

	reader := bufio.NewReader(os.Stdin)

	running := true
	for running {
		fmt.Print("SERVER>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Ошибка ввода:", err)
			break
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			color.Red("Exiting...")
			color.Red("No data was saved")
			running = false
		} else if input == "" {
			continue
		} else if input == "start server" {
			Variants()
		} else if input == "help" {
			instruction()
		}
	}
}
