package cli

import (
	"fmt"
	"sort"

	"github.com/RowMur/poke-cli/internal/cache"
)

func commandHelp(state *cliState, c *cache.CacheType, commandParams []string) error {
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