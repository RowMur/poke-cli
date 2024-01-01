package cli

import (
	"os"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/user"
)

func commandExit(state *user.CliState, c *cache.CacheType, commandParams []string) error {
	state.Save()
	os.Exit(0)
	return nil
}