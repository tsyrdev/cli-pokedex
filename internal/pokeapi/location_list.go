package pokeapi

import (
    "encoding/json"
    "net/http"
    "io"
)

func (c* Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    if val, ok := c.cache.Get(url); ok {
        locationsResp := RespShallowLocations{}
        if err := json.Unmarshal(val, &locationsResp); err != nil {
            return RespShallowLocations{}, err
        }

        return locationsResp, nil
    }

    req, err := http.NewRequest("GET", url, nil) 
    if err != nil {
        return RespShallowLocations{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowLocations{}, err
    }
    defer res.Body.Close()

    dat, err := io.ReadAll(res.Body)
    if err != nil {
        return RespShallowLocations{}, nil
    }

    var locationsResp RespShallowLocations
    if err = json.Unmarshal(dat, &locationsResp); err != nil {
        return RespShallowLocations{}, nil
    }

    c.cache.Add(url, dat)
    return locationsResp, nil
}
