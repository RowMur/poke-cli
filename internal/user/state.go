package user

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type CliState struct {
	Pokemap 		pokemap `json:"pokemap"`
	Pokedex         pokedex `json:"pokedex"`
	Mux             *sync.Mutex `json:"-"`
}

func (cs *CliState) Save() {
	_, stateFileLocation := getFileDetails()
	writeToFile(stateFileLocation, *cs)
}

func writeToFile(file string, state CliState) {
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
	initialState := CliState{
		Pokemap: pokemap{
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

func GetCliState() CliState {
	stateFileDir, stateFileLocation := getFileDetails()

	_, err := os.ReadFile(stateFileLocation)
	if err != nil {
		createInitialState(stateFileDir, stateFileLocation)
	}

	cs := CliState{}
	file, _ := os.ReadFile(stateFileLocation)
	json.Unmarshal(file, &cs)
	return cs
}