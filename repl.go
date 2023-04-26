package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func startRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Please enter some text: ")


		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0{
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
	
}

type cliCommand struct {
	name string
	description string  
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"map":{
			name: "map",
			description: "Lists of next page of location areas", 
			callback: callbackMap,
		},
		"mapb":{
			name: "map",
			description: "Lists the previous page of location areas", 
			callback: callbackMapb,
		},
		"exit": {
			name:"exit", 
			description: "Turns off the command",
			callback: callBackExit,
		},
	}
}

func cleanInput(str string) []string{
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}