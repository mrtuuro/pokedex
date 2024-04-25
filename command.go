package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	ExecuteFunc execFunc
}

type execFunc func(cfg *Config, fields []string) error

func NewCommand(name, desc string, callbackFunc execFunc) Command {
	return Command{
		Name:        name,
		Description: desc,
		ExecuteFunc: callbackFunc,
	}
}

func getCommandMap() map[string]Command {
	var commandMap = make(map[string]Command)
	commandMap["help"] = NewCommand("help", "Displays a help message.", commandHelp)
	commandMap["exit"] = NewCommand("exit", "Exit the CLI.", commandExit)
	commandMap["map"] = NewCommand("map", "Displays the names of next 20 locations.", commandMapForward)
	commandMap["mapb"] = NewCommand("mapb", "Displays the names of previous 20 locations.", commandMapBack)
	commandMap["explore"] = NewCommand("explore", "Displays all of the Pokemons in a given area.", commandExplore)

	return commandMap
}

func commandHelp(cfg *Config, fields []string) error {
	commandMap := getCommandMap()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands:")
	for _, v := range commandMap {
		fmt.Printf("  %v: %v\n", v.Name, v.Description)
	}
	return nil
}

func commandExit(cfg *Config, fields []string) error {
	os.Exit(1)
	return nil
}

func commandMapForward(cfg *Config, fields []string) error {
	locationResp, err := cfg.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locationResp.Next
	cfg.prev = locationResp.Previous

	for i, location := range locationResp.Results {

		fmt.Printf("%v. %v\n", i, location.Name)
	}
	return nil
}

func commandMapBack(cfg *Config, fields []string) error {
	locationResp, err := cfg.ListLocations(cfg.prev)
	if err != nil {
		return err
	}
	cfg.next = locationResp.Next
	cfg.prev = locationResp.Previous

	for i, location := range locationResp.Results {

		fmt.Printf("%v. %v\n", i, location.Name)
	}
	return nil
}

func commandExplore(cfg *Config, fields []string) error {
    param := fields[1]
    location, err := cfg.ListPokemons(cfg.next, param)
    if err != nil {
        return err
    }
    for _, enc := range location.PokemonEncounters {
        fmt.Printf("%v\n", enc.Pokemon.Name)
    }
	return nil
}




