package main

import "fmt"

func commandHelp(c *config, args ...string) error {
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range GetCommands() {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
    fmt.Println()
    return nil
}
