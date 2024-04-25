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

type execFunc func(cfg *Config) error

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

	return commandMap
}

func commandHelp(cfg *Config) error {
	commandMap := getCommandMap()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands:")
	for _, v := range commandMap {
		fmt.Printf("  %v: %v\n", v.Name, v.Description)
	}
	return nil
}

func commandExit(cfg *Config) error {
	os.Exit(1)
	return nil
}

func commandMapForward(cfg *Config) error {
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

func commandMapBack(cfg *Config) error {
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
