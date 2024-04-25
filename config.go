package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mrtuuro/pokedex/internal/cache"
)

type Config struct {
	BaseURL string
	Client  http.Client
	next    *string
	prev    *string

	Cache *cache.Cache
}

var baseUrl = "https://pokeapi.co/api/v2"

func NewConfig() *Config {

	newClient := http.Client{
		Timeout: 20 * time.Second,
	}
	return &Config{
		BaseURL: baseUrl,
		Client:  newClient,
		Cache:   cache.NewCache(5 * time.Minute),
	}
}

func (c *Config) ListLocations(pageURL *string) (RespLocation, error) {
	url := c.BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.Cache.Get(url); ok {
		locationsResp := RespLocation{}
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return RespLocation{}, err
		}
        fmt.Println("this is from cache")
		return locationsResp, nil
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

    fmt.Println("this is from request")
	c.Cache.Add(url, dat)
	return locationsResp, nil

}
