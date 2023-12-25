package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*CliState, *cacheType, []string) error
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
	}
}

func commandHelp(state *CliState, cache *cacheType, commandParams []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Poke CLI!")
	fmt.Println("Usage:")
	fmt.Println()
	
	cliCommands := getCliCommands()

	keys := make([]string, 0)
	for key := range cliCommands {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %s\n", cliCommands[key].name, cliCommands[key].description)
	}

	fmt.Println()
	return nil
}

func commandExit(state  *CliState, cache *cacheType, commandParams []string) error {
	os.Exit(0)
	return nil
}

func genericMapCommand(url string, state *CliState, cache *cacheType) error {
	locations, err := fetchLocations(url, cache)
	if err != nil {
		fmt.Println(err)
		return err
	}

	state.prevLocationURL = &locations.Previous
	state.nextLocationURL = &locations.Next

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
} 

func commandMap(state *CliState, cache *cacheType, commandParams []string) error {
	return genericMapCommand(*state.nextLocationURL, state, cache)
}

func commandMapBack(state *CliState, cache *cacheType, commandParams []string) error {
	return genericMapCommand(*state.prevLocationURL, state, cache)
}

func commandExplore(state *CliState, cache *cacheType, commandParams []string) error {
	areaName := commandParams[0]
	fmt.Printf("Exploring %s\n", areaName)

	location, err := fetchLocationDetail(areaName, cache)
	if err != nil {
		if err.Error() != "404" {
			fmt.Println(err)
			return err
		}

		fmt.Println("...I think you got lost. Check your spelling.")
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

func commandCatch(state *CliState, cache *cacheType, commandParams []string) error {
	pokemonToCatchName := commandParams[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatchName)

	p, err := fetchPokemon(pokemonToCatchName, cache)
	if err != nil {
		if err.Error() != "404" {
			fmt.Println(err)
			return err
		}

		fmt.Println("...pokeball missed. Check your spelling.")
		return err
	}

	wasCaught := rand.Intn(256) >= p.BaseExperience
	if !wasCaught {
		fmt.Printf("%s escaped!\n", pokemonToCatchName)
		return nil
	}

	prevPokedex := *state.pokedex
	_, ok := prevPokedex[p.Name]
	if !ok {
		prevPokedex[p.Name] = PokedexEntry{
			pokemon: p,
			timesCaught: 1,
		}
	} else {
		prevTimesCaught := prevPokedex[p.Name].timesCaught
		prevPokedex[p.Name] = PokedexEntry{
			pokemon: p,
			timesCaught: prevTimesCaught + 1,
		}
	}

	state.mux.Lock()
	defer state.mux.Unlock()
	state.pokedex = &prevPokedex

	pokedex := *state.pokedex
	entry := pokedex[p.Name]
	fmt.Printf("%s was caught! Caught a total of %v times\n", entry.pokemon.Name, entry.timesCaught)
	return nil
}

func commandInspect(state *CliState, cache *cacheType, commandParams []string) error {
	pokemonToInspect := commandParams[0]
	pokedex := *state.pokedex

	pokedexEntry, ok := pokedex[pokemonToInspect]
	if !ok {
		fmt.Printf("you have not caught %s yet!\n", pokemonToInspect)
		return nil
	}

	pokemon := pokedexEntry.pokemon
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemonType.Type.Name)
	}

	return nil
}

type PokedexEntry struct {
	pokemon Pokemon
	timesCaught int
}

type CliState struct {
	prevLocationURL *string
	nextLocationURL *string
	pokedex *map[string]PokedexEntry
	mux *sync.Mutex
}

func Cli () {
	cliCommands := getCliCommands()

	var initLocationURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	CliState := CliState{
		prevLocationURL: &initLocationURL,
		nextLocationURL: &initLocationURL,
		pokedex: &map[string]PokedexEntry{},
		mux: &sync.Mutex{},
	}

	cache := newCache(time.Duration(5) * time.Second)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	commandHelp(&CliState, &cache, []string{})

	for {
		fmt.Print("Poke CLI> ")
		scanner.Scan()
		enteredText := scanner.Text()
		enteredFields := strings.Fields(enteredText)
		
		enteredCommand := enteredFields[0]
		enteredParameters := enteredFields[1:]

		command, ok := cliCommands[enteredCommand]
		if !ok {
			fmt.Printf("'%s' is an invalid command\nUse 'help' to see a list of commands\n", enteredCommand)
			continue
		}

		command.callback(&CliState, &cache, enteredParameters)
	}
}