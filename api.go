package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func fetch(url string) ([]byte, error) {
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

func fetchLocations(url string) (locationResponse, error) {
	body, err := fetch(url)
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
