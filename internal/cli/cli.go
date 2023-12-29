package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/RowMur/poke-cli/internal/cache"
	"github.com/RowMur/poke-cli/internal/pokedata"
)

type pokedexEntry struct {
	pokemon pokedata.Pokemon
	timesCaught int
}

type cliState struct {
	prevLocationURL *string
	nextLocationURL *string
	pokedex *map[string]pokedexEntry
	mux *sync.Mutex
}

func Cli () {
	cliCommands := getCliCommands()

	var initLocationURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	CliState := cliState{
		prevLocationURL: &initLocationURL,
		nextLocationURL: &initLocationURL,
		pokedex: &map[string]pokedexEntry{},
		mux: &sync.Mutex{},
	}

	c := cache.NewCache(time.Duration(5) * time.Second)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	commandHelp(&CliState, &c, []string{})

	for {
		fmt.Print("Poke CLI> ")
		scanner.Scan()
		enteredText := scanner.Text()
		enteredFields := strings.Fields(enteredText)
		
		enteredCommand := enteredFields[0]
		enteredParameters := enteredFields[1:]

		command, ok := cliCommands[enteredCommand]
		if !ok {
			fmt.Printf("'%s' is an invalid command\nUse 'help' to see a list of commands\n", enteredCommand)
			continue
		}

		command.callback(&CliState, &c, enteredParameters)
	}
}