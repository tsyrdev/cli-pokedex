package repl 

import (
    "fmt"
    "strings" 
    "os"
    "bufio"
    "github.com/tsyrdev/pokedex/internal/commands"
)


func Start() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        in := scanner.Text() 
        cleanIn := cleanInput(in)

        commandName := cleanIn[0]
        command, exists := commands.GetCommands()[commandName]
        if exists {
            err := command.Callback()
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("Unknown command")
            continue
        }
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

