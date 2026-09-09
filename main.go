package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	textSlice := strings.Fields(strings.ToLower((text)))
	return textSlice
}

func main() {
	/*inputString := strings.Join(os.Args[1:], " ")
	fmt.Println("os.Args input was: ", inputString)
	fmt.Println("cleanInput: ", cleanInput(inputString))*/
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanToken := cleanInput(scanner.Text())
		if len(cleanToken) == 0 {
			continue
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input: %s", err)
		}
		if com, ok := configCom.commands[cleanToken[0]]; ok {
			err := com.callback(&configCom)
			if err != nil {
				fmt.Println("Error: %w", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
