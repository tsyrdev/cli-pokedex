package pokeapi

import (
    "encoding/json"
    "net/http"
    "io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
    url := baseURL + "/pokemon/" + pokemonName

    if bytes, ok := c.cache.Get(url); ok {
        var pokemon Pokemon
        if err := json.Unmarshal(bytes, &pokemon); err != nil {
            return Pokemon{}, err
        }

        return pokemon, nil
    }
     
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return Pokemon{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }
    defer res.Body.Close()

    dat, err := io.ReadAll(res.Body)
    if err != nil {
        return Pokemon{}, err 
    }

    var pokemon Pokemon
    if err := json.Unmarshal(dat, &pokemon); err != nil {
        return Pokemon{}, err 
    }

    c.cache.Add(url, dat)
    return pokemon, nil 
}
