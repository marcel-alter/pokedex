package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func initCommands() {

	configCom.commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	configCom.commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	configCom.commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas of the next page",
		callback:    commandMap,
	}
	configCom.commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of 20 location areas of the previous page",
		callback:    commandMapB,
	}
}

// Command functions ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	helpText := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, comm := range configCom.commands {
		//iterate through map values
		helpText = helpText + fmt.Sprintf("%s: %s\n", comm.name, comm.description)

	}
	fmt.Print(helpText)
	return nil
}

// Command MAP FUNCTIONS ↓↓↓↓↓↓↓
func mapStart() error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return fmt.Errorf("Error requesting maps: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error fetching maps: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll failed: %w", err)
	}
	if err := json.Unmarshal(body, &configCom.location); err != nil {
		return fmt.Errorf("Error Unmarshaling maps: %w", err)
	}
	configCom.location.PAGE += 1
	fmt.Printf("PAGE: %d ↓↓↓↓↓↓\n", configCom.location.PAGE)
	for _, area := range configCom.location.RESULTS {
		fmt.Print(area.NAME + "\n")
	}
	fmt.Printf("PAGE: %d ↑↑↑↑↑↑\n", configCom.location.PAGE)
	return nil
}

func mapNext() error {
	nextURL := configCom.location.NEXT
	client := &http.Client{}
	req, err := http.NewRequest("GET", nextURL, nil)
	if err != nil {
		return fmt.Errorf("Error requesting maps: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error fetching maps: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll failed: %w", err)
	}
	if err := json.Unmarshal(body, &configCom.location); err != nil {
		return fmt.Errorf("Error Unmarshaling maps: %w", err)
	}
	configCom.location.PAGE += 1
	fmt.Printf("PAGE: %d ↓↓↓↓↓↓\n", configCom.location.PAGE)
	for _, area := range configCom.location.RESULTS {
		fmt.Print(area.NAME + "\n")
	}
	fmt.Printf("PAGE: %d ↑↑↑↑↑↑\n", configCom.location.PAGE)
	return nil
}

func mapBack() error {
	previousURL := configCom.location.PREVIOUS
	client := &http.Client{}
	req, err := http.NewRequest("GET", previousURL, nil)
	if err != nil {
		return fmt.Errorf("Error requesting maps: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error fetching maps: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll failed: %w", err)
	}
	if err := json.Unmarshal(body, &configCom.location); err != nil {
		return fmt.Errorf("Error Unmarshaling maps: %w", err)
	}
	configCom.location.PAGE -= 1
	fmt.Printf("PAGE: %d ↓↓↓↓↓↓\n", configCom.location.PAGE)
	for _, area := range configCom.location.RESULTS {
		fmt.Print(area.NAME + "\n")
	}
	fmt.Printf("PAGE: %d ↑↑↑↑↑↑\n", configCom.location.PAGE)
	return nil
}

func commandMap(*config) error {
	if configCom.location.PAGE != 0 {
		return mapNext()
	}
	return mapStart()
}

func commandMapB(*config) error {
	if configCom.location.PAGE <= 1 {
		fmt.Print("Use command 'map' first\n")
		return nil
	}
	return mapBack()
}

// MAP FUNCTIONS ↑↑↑↑↑↑↑↑↑↑↑
