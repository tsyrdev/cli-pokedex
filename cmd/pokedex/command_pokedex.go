package main 

import (
    "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    if len(args) != 0 {
        return fmt.Errorf("No arguments may be provided with this command")
    }
    
    fmt.Println("Your Pokedex:")
    for _, v := range cfg.caughtPokemon {
        fmt.Println(" -", v.Name)
    }
    return nil
}
