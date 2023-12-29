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

	prevPokedex := *state.pokedex
	_, ok := prevPokedex[p.Name]
	if !ok {
		prevPokedex[p.Name] = pokedexEntry{
			pokemon:     p,
			timesCaught: 1,
		}
	} else {
		prevTimesCaught := prevPokedex[p.Name].timesCaught
		prevPokedex[p.Name] = pokedexEntry{
			pokemon:     p,
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