package main

import (
    "fmt"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return fmt.Errorf("Error executing explore command: Only 1 location argument may be provided")
    }
    
    location, err := cfg.pokeapiClient.GetLocation(args[0])
    if err != nil {
        return err
    }
    fmt.Printf("Exploring %s...\n", location.Name)
    fmt.Println("Found Pokemon:")
    for _, encounter := range location.PokemonEncounters {
        fmt.Printf(" - %s\n", encounter.Pokemon.Name)
    }
    return nil
}
    
