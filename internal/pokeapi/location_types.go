package pokeapi

type RespShallowLocations struct {
    Count       int     `json:"count"`
    Next        *string `json:"next"`
    Previous    *string `json:"previous"`
    Results     []struct {
        Name    string `json:"name"`
        URL     string `json:"url"`
    } `json:"results"`
}

type Location struct {
    EncounterMethodRates    []encounterMethodRates  `json:"encounter_method_rates"`
    GameIndex               int                     `json:"game_index"`
    Id                      int                     `json:"id"`
    Location                nameURLPair             `json:"location"`
    Name                    string                  `json:"name"`
    Names                   []name                  `json:"names"`
    PokemonEncounters       []encounter             `json:"pokemon_encounters"`
}

type encounter struct {
    Pokemon         nameURLPair         `json:"pokemon"`
    VersionDetails  []encounterVersion  `json:"version_details"`
}

type encounterVersion struct {
    EncounterDetails    []encounterDetails  `json:"encounter_details"`
    MaxChance           int                 `json:"max_chance"`
    Version             nameURLPair         `json:"version"`
}

type encounterDetails struct {
    Chance              int             `json:"chance"`
    ConditionValues     []interface{}   `json:"condition_values"`
    MaxLevel            int             `json:"max_level"` 
    Method              nameURLPair     `json:"method"`
    MinLevel            int             `json:"min_level"`
}

type name struct {
    Language    nameURLPair `json:"language"`
    Name        string      `json:"name"`
}

type encounterMethodRates struct {
    EncounterMethod nameURLPair         `json:"encounter_method"`
    VersionDetails  []versionDetails    `json:"versions_details"`
}

type versionDetails struct {
    Rate    int         `json:"rate"`
    Version nameURLPair `json:"version"`
}

type nameURLPair struct {
    Name    string  `json:"name"`
    URL     string  `json:"url"`
}
