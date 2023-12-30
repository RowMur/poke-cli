package cli

import (
	"fmt"
	"math/rand"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/pokedata"
)

func commandCatch(state *cliState, c *cache.CacheType, commandParams []string) error {
	pokemonToCatchName := commandParams[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatchName)

	p, err := pokedata.FetchPokemon(pokemonToCatchName, c)
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

	state.Pokedex.addEntry(p)

	entry := state.Pokedex.Entries[p.Name]
	fmt.Printf("%s was caught! Caught a total of %v times\n", entry.Pokemon.Name, entry.TimesCaught)
	return nil
}