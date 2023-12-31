package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/user"
)

func commandInspect(state *user.CliState, c *cache.CacheType, commandParams []string) error {
	pokemonToInspect := commandParams[0]
	pokedex := state.Pokedex

	pokedexEntry, ok := pokedex.Entries[pokemonToInspect]
	if !ok {
		fmt.Printf("you have not caught %s yet!\n", pokemonToInspect)
		return nil
	}

	pokemon := pokedexEntry.Pokemon
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