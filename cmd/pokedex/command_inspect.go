package main

import (
    "fmt"
)

func commandInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
        return fmt.Errorf("Error inspecting pokemon: Should take exactly 1 parameter")
    }
    
    pokemon, ok := cfg.caughtPokemon[args[0]]
    if !ok {
        fmt.Printf("you have not caught that pokemon\n")
        return nil
    }

    fmt.Printf(`
Name: %s 
Height: %d 
Weight: %d 
`, pokemon.Name, pokemon.Height, pokemon.Weight)

    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("   -%s: %v\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Println("   -", typeInfo.Type.Name)
    }
    return nil
}
