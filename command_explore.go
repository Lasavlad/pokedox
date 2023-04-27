package main

import (
	"fmt" 
	"errors"
)
 

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No location areas")
	}
	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil{
		return err
	}
	fmt.Printf("Pokemon in %s", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters{
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	
	return nil
}

