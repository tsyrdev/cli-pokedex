package pokeapi

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
    url := baseURL + "/location-area/" + locationName

    if val, ok := c.cache.Get(url); ok {
        location := Location{}
        if err := json.Unmarshal(val, &location); err != nil {
            return Location{}, err
        }
        
        return location, nil
    }

    req, err := http.NewRequest("GET", url, nil) 
    if err != nil {
        return Location{}, fmt.Errorf("Error creating the pokemon list request: %w", err)
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return Location{}, fmt.Errorf("Error making the pokemon list request: %w", err)
    }
    defer res.Body.Close()

    dat, err := io.ReadAll(res.Body)
    if err != nil {
        return Location{}, err
    }
    
    var location Location
    if err = json.Unmarshal(dat, &location); err != nil {
        return Location{}, fmt.Errorf("Error decoding the pokemon list response: %w", err)
    }

    c.cache.Add(url, dat)
    return location, nil
}

