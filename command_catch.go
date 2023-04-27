package main

import (
	"fmt" 
	"errors"
	"math/rand"
)
 

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No Pokemon name provided ")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil{
		return err
	}
	const threshold = 50

	randNum := rand.Intn(pokemon.BaseExperience)
	
	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s:", pokemonName)
	}
	
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

