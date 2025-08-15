package pokeapi

// This is the type struct needed to decode the
// JSON data from the locations url

type DecLocations struct {
	Count    int       `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []cityloc `json:"results"`
}

type cityloc struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
