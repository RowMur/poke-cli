package main

import (
	"bufio"
	"fmt"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*CliConfig) error
}

func getCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"map": {
			name: "map",
			description: "Display a list of locations",
			callback: commandMap,
		},
		"mapBack": {
			name: "mapBack",
			description: "Display the previous list of locations",
			callback: commandMapBack,
		},
	}
}

func commandHelp(config *CliConfig) error {
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

func commandExit(config *CliConfig) error {
	os.Exit(0)
	return nil
}

func genericMapCommand(url string, config *CliConfig) error {
	locations, err := fetchLocations(url)
	if err != nil {
		fmt.Println(error.Error(err))
		return err
	}

	config.prevLocationURL = &locations.Previous
	config.nextLocationURL = &locations.Next

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
} 

func commandMap(config *CliConfig) error {
	return genericMapCommand(*config.nextLocationURL, config)
}

func commandMapBack(config *CliConfig) error {
	return genericMapCommand(*config.prevLocationURL, config)
}

type CliConfig struct {
	prevLocationURL *string
	nextLocationURL *string
}

func Cli () {
	cliCommands := getCliCommands()

	var initLocationURL string = "https://pokeapi.co/api/v2/location"
	cliConfig := CliConfig{
		prevLocationURL: &initLocationURL,
		nextLocationURL: &initLocationURL,
	}

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("Poke CLI> ")
		scanner.Scan()
		enteredCommand := scanner.Text()

		command, ok := cliCommands[enteredCommand]
		if !ok {
			fmt.Printf("'%s' is an invalid command\n", enteredCommand)
			continue
		}

		command.callback(&cliConfig)
	}
}