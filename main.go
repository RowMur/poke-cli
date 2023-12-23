package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Poke CLI",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Poke CLI!")
	fmt.Println("Usage:")
	fmt.Println()
	
	cliCommands := getCliCommands()
	for _, value := range cliCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	cliCommands := getCliCommands()

	for {
		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)
		
		fmt.Print("Poke CLI > ")
		scanner.Scan()

		enteredCommand := scanner.Text()
		command, ok := cliCommands[enteredCommand]
		if !ok {
			fmt.Printf("'%s' is an invalid command\n", enteredCommand)
			continue
		}

		command.callback()
	}
}
