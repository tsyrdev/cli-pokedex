package repl 

import (
    "fmt"
    "strings" 
    "os"
    "bufio"
)
func Start() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        in := scanner.Text() 
        cleanIn := cleanInput(in)
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}
