package cli

import (
	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/user"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*user.CliState, *cache.CacheType, []string) error
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
		"catch": {
			name: "catch",
			description: "Attempts to catch a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspects a pokemon from the pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists Pokedex entries",
			callback: commandPokedex,
		},
	}
}