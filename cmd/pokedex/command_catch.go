package main 

import (
    "fmt"
    "math/rand"
)

func commandCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return fmt.Errorf("Error while catching pokemon: Exactly 1 argument must be provided")
    }

    pokemon, err := cfg.pokeapiClient.GetPokemon(args[0]) 
    if err != nil {
        return err
    }

    res := rand.Intn(pokemon.BaseExperience)

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
    if res > 40 {
        fmt.Printf("%s escaped!\n", pokemon.Name)
        return nil
    }

    fmt.Printf("%s has been captured!\n", pokemon.Name)
    fmt.Println("You may now inspect it with the inspect command.")

    cfg.caughtPokemon[pokemon.Name] = pokemon
    return nil
}
