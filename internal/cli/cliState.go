package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type cliState struct {
	MapState mapState `json:"mapState"`
	Pokedex         pokedex `json:"pokedex"`
	Mux             *sync.Mutex `json:"-"`
}

func (cs *cliState) save() {
	_, stateFileLocation := getFileDetails()
	writeToFile(stateFileLocation, *cs)
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

func createInitialState(dir, file string) {
	initialLocationURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	initialState := cliState{
		MapState: mapState{
			PrevLocationURL: initialLocationURL,
			NextLocationURL: initialLocationURL,
		},
		Pokedex: pokedex{
			Entries: map[string]pokedexEntry{},
		},
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