package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
	Client  http.Client
	next    *string
	prev    *string
}

var baseUrl = "https://pokeapi.co/api/v2"

func NewConfig() *Config {

	newClient := http.Client{
		Timeout: 20 * time.Second,
	}
	return &Config{
		BaseURL: baseUrl,
		Client:  newClient,
	}
}

func (c *Config) ListLocations(pageURL *string) (RespLocation, error) {
	url := c.BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationsResp := RespLocation{}
	if err = json.Unmarshal(dat, &locationsResp); err != nil {
		return RespLocation{}, err
	}
	return locationsResp, nil

}
