# Cli-Pokedex: A cli-tool that interacts with the Pokeapi
![Static Badge](https://img.shields.io/badge/License-MIT-red.svg)
## Overview

The cli-pokedex tool allows the user to look-up data on different pokemons and pokemon locations, and also has a few game-like features. 
It is written in Go due to its simplicity and HTTP package features, which make it extremely easily to use this sort of API. 
You can fork the project and add your own extra commands (or change existing ones) if you'd like. 

## How to Install and Run 

### Prerequisites
- Have Go installed 
### Installation Steps
1. Clone the repo
```
git clone https://github.com/tsyrdev/cli-pokedex.git
cd pokedex
```
2. Build the project 
```
go build -o cliPokedex ./cmd/cli-pokedex/
```

You're good to go! 

## How to Use it 

Currently there are 7 commands able. 
You can check what they do and how to use them using the `help` command: 
```
Pokedex > help

Welcome to the Pokedex!
Usage:

mapb: Get the previous page of locations
explore <location_name>: Explore a location
catch <pokemon_name>: Try to catch a pokemon
inspect <pokemon_name>: Inspect one of the pokemons you've caught
pokedex: Display your pokemons
exit: Exit the Pokedex
help: Get help
map: Get the next page of locations
```

## Credits 

This project idea comes from a Go tutorial on [Boot.dev](https://www.boot.dev/tracks/backend)

## License 

This project is licensed under the [MIT License](LICENSE)
