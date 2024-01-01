package cli

import (
	"fmt"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/pokedata"
	"github.com/RowMur/poke-cli/internal/user"
)

func commandExplore(state *user.CliState, c *cache.CacheType, commandParams []string) error {
	areaName := commandParams[0]
	fmt.Printf("Exploring %s\n", areaName)

	location, err := pokedata.FetchLocationDetail(areaName, c)
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