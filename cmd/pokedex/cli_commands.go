package main

type cliCommand struct {
    Name        string
    Description string
    Callback    func(*config, ...string) error
}

func GetCommands() map[string]cliCommand{
    return map[string]cliCommand {
        "exit": {
            Name:           "exit",
            Description:    "Exit the Pokedex",
            Callback:       commandExit,
        },
        "help": {
            Name:           "help",
            Description:    "Get help",
            Callback:       commandHelp,
        },
        "map": {
            Name:           "map", 
            Description:    "Get the next page of locations",
            Callback:       commandMapf,
        },
        "mapb": {
            Name:           "mapb",
            Description:    "Get the previous page of locations",
            Callback:       commandMapb,
        },
        "explore": {
            Name:           "explore <location_name>",
            Description:    "Explore a location",
            Callback:       commandExplore,
        },
        "catch": {
            Name:           "catch <pokemon_name>",
            Description:    "Try to catch a pokemon",
            Callback:       commandCatch,
        },
    }
}
