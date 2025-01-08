package pokeapi

type Config struct {
	Page    int
	Results *[]Location
	//AreaId   map[string]int
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
}

type APIResponse struct {
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
	Pokemons []Pokemons `json:"pokemon_encounters"`
}

type Pokemons struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func InitConf() *Config {
	return &Config{
		Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}
}

func (c *Config) Update(newNext string, newPrev *string) {
	c.Next, c.Previous = newNext, newPrev
}
func (c *Config) PageCounter(input []string) error {
	if len(input) > 1 {
		return nil
	}
	switch input[0] {
	case "map":
		c.Page++
		return nil
	case "mapb":
		if c.Page != 1 && c.Page > 0 {
			c.Page--
		}
		return nil
	default:
		return nil
	}
}

/*func (c *Config) TakeResult(api *APIResponse) {
	c.AreaId = make(map[string]int, 20)
	c.Results = &api.Results

	for id, loc := range *c.Results {
		c.AreaId[loc.Name] = id
	}
}*/
