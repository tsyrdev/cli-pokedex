package main 

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        in := scanner.Text() 
        cleanIn := cleanInput(in)
        fmt.Printf("Your command was: %s\n", cleanIn[0])
    }
}
