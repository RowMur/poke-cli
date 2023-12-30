package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
)

func commandPokedex(state *cliState, c *cache.CacheType, commandParams []string) error {
	pokedex := state.Pokedex

	fmt.Printf("Your Pokedex:\n")
	for key := range pokedex.Entries {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}