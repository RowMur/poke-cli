package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*CliConfig, *cacheType, []string) error
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
		"explore": {
			name: "explore",
			description: "Explores a given area and prints list of found Pokemon",
			callback: commandExplore,
		},
	}
}

func commandHelp(config *CliConfig, cache *cacheType, commandParams []string) error {
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

func commandExit(config *CliConfig, cache *cacheType, commandParams []string) error {
	os.Exit(0)
	return nil
}

func genericMapCommand(url string, config *CliConfig, cache *cacheType) error {
	locations, err := fetchLocations(url, cache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	config.prevLocationURL = &locations.Previous
	config.nextLocationURL = &locations.Next

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
} 

func commandMap(config *CliConfig, cache *cacheType, commandParams []string) error {
	return genericMapCommand(*config.nextLocationURL, config, cache)
}

func commandMapBack(config *CliConfig, cache *cacheType, commandParams []string) error {
	return genericMapCommand(*config.prevLocationURL, config, cache)
}

func commandExplore(config *CliConfig, cache *cacheType, commandParams []string) error {
	areaName := commandParams[0]
	fmt.Printf("Exploring %s\n", areaName)

	location, err := fetchLocationDetail(areaName, cache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	encounters := location.PokemonEncounters
	if len(encounters) == 0 {
		fmt.Printf("Found no Pokemon in area: %s\n", areaName)
		return nil
	}

	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

type CliConfig struct {
	prevLocationURL *string
	nextLocationURL *string
}

func Cli () {
	cliCommands := getCliCommands()

	var initLocationURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cliConfig := CliConfig{
		prevLocationURL: &initLocationURL,
		nextLocationURL: &initLocationURL,
	}

	cache := newCache(time.Duration(5) * time.Second)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("Poke CLI> ")
		scanner.Scan()
		enteredText := scanner.Text()
		enteredFields := strings.Fields(enteredText)
		
		enteredCommand := enteredFields[0]
		enteredParameters := enteredFields[1:]

		command, ok := cliCommands[enteredCommand]
		if !ok {
			fmt.Printf("'%s' is an invalid command\n", enteredCommand)
			continue
		}

		command.callback(&cliConfig, &cache, enteredParameters)
	}
}