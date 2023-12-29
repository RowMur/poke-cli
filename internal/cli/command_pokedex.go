package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
)

func commandPokedex(state *cliState, c *cache.CacheType, commandParams []string) error {
	pokedex := *state.pokedex

	fmt.Printf("Your Pokedex:\n")
	for key := range pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}