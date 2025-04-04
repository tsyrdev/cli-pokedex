# Cli-Pokedex: A CLI Tool That Interacts With the PokéAPI
![License: MIT](https://img.shields.io/badge/License-MIT-red.svg)
## Overview

The **cli-pokedex** tool allows the user to look up data on different Pokémon and locations, with a few game-like features. 
It is written in **Go** due to its **simplicity, efficiency, and built-in support for handling HTTP requests**. Go's **concurrency model** also ensures smooth and fast execution,
even when making multiple API calls.

Additionally, Go's **lightweight binaries** and **cross-platform compatibility** make it an excellent choice for a CLI-based application,
enabling **fast performance** and **easy distribution**.

You can **fork** the project and add custom commands or modify existing ones as needed.

## How to Install and Run 

### Prerequisites
- Go must be installed (download it from [golang.org](https://go.dev/dl/)).
- Ensure your `$GOPATH/bin` is in your system's `PATH`.
  
### **Option 1: Install the binary (Recommended)**
Run the following command to install the binary globally: 
```sh
go install github.com/tsyrdev/cli-pokedex/cmd/pokedex@latest
```
Once installed, you can run the tool using `pokedex`.

### **Option 2: Build from source**
1. Clone the repo:
```sh
git clone https://github.com/tsyrdev/cli-pokedex.git
cd cli-pokedex
```
2. Build the project:
```sh
go build -o pokedex ./cmd/cli-pokedex/
```
3. Run the tool locally:
```sh
./pokedex
```
4. If you want to run it globally, move the binary to your `$GOPATH/bin`:
```sh
mv pokedex ~/go/bin
```

## How to Use it 

The `cli-pokedex` tool supports **7 commands**. You can see available commands by using the the `help` command: 
```
Pokedex > help

Welcome to the Pokedex!
Usage:

mapb: Get the previous page of locations
explore <location_name>: Explore a location
catch <pokemon_name>: Try to catch a pokemon
inspect <pokemon_name>: Inspect one of the pokemons you've caught
pokedex: Display your Pokémon
exit: Exit the Pokedex
help: Get help
map: Get the next page of locations
```

## Credits 

This project idea comes from a Go tutorial on [Boot.dev](https://www.boot.dev/tracks/backend)

## License 

This project is licensed under the [MIT License](LICENSE)
