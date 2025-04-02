package pokeapi

type Config struct {
    pokeapiClient   Client
    nextLocationURL *string
    prevLocationURL *string
}
