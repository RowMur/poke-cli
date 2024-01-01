package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/pokedata"
	"github.com/RowMur/poke-cli/internal/user"
)

func genericMapCommand(url string, state *user.CliState, c *cache.CacheType) error {
	locations, err := pokedata.FetchLocations(url, c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	state.Pokemap.UpdateURLs(locations.Previous, locations.Next)

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMap(state *user.CliState, c *cache.CacheType, commandParams []string) error {
	return genericMapCommand(state.Pokemap.NextLocationURL, state, c)
}

func commandMapBack(state *user.CliState, c *cache.CacheType, commandParams []string) error {
	return genericMapCommand(state.Pokemap.PrevLocationURL, state, c)
}