package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/RowMur/poke-cli/internal/pokedata"
)

type pokedexEntry struct {
	Pokemon     pokedata.Pokemon `json:"pokemon"`
	TimesCaught int `json:"timesCaught"`
}

type cliState struct {
	PrevLocationURL string `json:"prevLocationURL"`
	NextLocationURL string `json:"nextLocationURL"`
	Pokedex         map[string]pokedexEntry `json:"pokedex"`
	Mux             *sync.Mutex `json:"-"`
}

func writeToFile(file string, state cliState) {
	stateData, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile(file, stateData, 0644)
}

func getFileDetails() (fileDir, fileLocation string) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Printf("error getting users cache directory: %s\n", err)
	}

	stateFileDir := fmt.Sprintf("%s\\poke-cli", cacheDir)
	stateFileLocation := fmt.Sprintf("%s\\cliState.json", stateFileDir)

	return stateFileDir, stateFileLocation
}

func (cs *cliState) Update(newState cliState) {
	_, stateFileLocation := getFileDetails()
	cs = &newState
	writeToFile(stateFileLocation, newState)
}

func createInitialState(dir, file string) {
	initialLocationURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	initialState := cliState{
		PrevLocationURL: initialLocationURL,
		NextLocationURL: initialLocationURL,
		Pokedex: map[string]pokedexEntry{},
		Mux: &sync.Mutex{},
	}

	os.Mkdir(dir, os.ModePerm)
	writeToFile(file, initialState)	
}

func getCliState() cliState {
	stateFileDir, stateFileLocation := getFileDetails()

	_, err := os.ReadFile(stateFileLocation)
	if err != nil {
		createInitialState(stateFileDir, stateFileLocation)
	}

	cs := cliState{}
	file, _ := os.ReadFile(stateFileLocation)
	json.Unmarshal(file, &cs)
	return cs
}