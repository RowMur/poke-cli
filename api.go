package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func fetch(url string, cache *cacheType) ([]byte, error) {
	cachedEntry, ok := cache.get(url)
	if ok {
		return cachedEntry, nil
	}

	res, err := http.Get(url)
	if err != nil {
		errorMessage := fmt.Sprintf("error fetching url: '%s'", url)
		return []byte{}, errors.New(errorMessage)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		errorMessage := fmt.Sprintf("error with status code: '%v'", res.StatusCode)
		return []byte{}, errors.New(errorMessage)
	}

	cache.add(url, body)
	return body, nil
}

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchLocations(url string, cache *cacheType) (locationResponse, error) {
	body, err := fetch(url, cache)
	if err != nil {
		return locationResponse{}, err
	}

	data := locationResponse{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return locationResponse{}, err
	}

	return data, nil
}

type locationDetailResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func fetchLocationDetail(location string, cache *cacheType) (locationDetailResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	body, err := fetch(url, cache)
	if err != nil {
		return locationDetailResponse{}, err
	}

	data := locationDetailResponse{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		return locationDetailResponse{}, err
	}

	return data, nil
}
