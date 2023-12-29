package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/pokedata"
)

func genericMapCommand(url string, state *cliState, c *cache.CacheType) error {
	locations, err := pokedata.FetchLocations(url, c)
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

func commandMap(state *cliState, c *cache.CacheType, commandParams []string) error {
	return genericMapCommand(*state.nextLocationURL, state, c)
}

func commandMapBack(state *cliState, c *cache.CacheType, commandParams []string) error {
	return genericMapCommand(*state.prevLocationURL, state, c)
}