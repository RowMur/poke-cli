package cli

import "github.com/RowMur/poke-cli/internal/pokedata"

type pokedexEntry struct {
	Pokemon     pokedata.Pokemon `json:"pokemon"`
	TimesCaught int              `json:"timesCaught"`
}

type pokedex struct {
	Entries map[string]pokedexEntry `json:"entries"`
}

func (pd *pokedex) addEntry(pokemon pokedata.Pokemon) {
	entry, ok := pd.Entries[pokemon.Name]
	if ok {
		entry.TimesCaught += 1
		return
	}

	pd.Entries[pokemon.Name] = pokedexEntry{
		Pokemon: pokemon,
		TimesCaught: 1,
	}
}
