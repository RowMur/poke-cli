package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RowMur/poke-cli/internal/cache"
)

func Cli () {
	cliCommands := getCliCommands()
	c := cache.NewCache(time.Duration(5) * time.Second)
	cs := getCliState()

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	commandHelp(&cs, &c, []string{})

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

		command.callback(&cs, &c, enteredParameters)
	}
}