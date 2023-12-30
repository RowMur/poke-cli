package cli

import (
	"os"

	"github.com/RowMur/poke-cli/internal/cache"
)

func commandExit(state *cliState, c *cache.CacheType, commandParams []string) error {
	state.save()
	os.Exit(0)
	return nil
}