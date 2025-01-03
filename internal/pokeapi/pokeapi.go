package pokeapi

type Config struct {
	Page     int
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
}

type APIResponse struct {
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
}

func InitConf() *Config {
	return &Config{
		Next: "https://pokeapi.co/api/v2/location-area",
	}
}

func (c *Config) Update(newNext string, newPrev *string) {
	c.Next, c.Previous = newNext, newPrev
}
func (c *Config) PageCounter(command string) error {
	switch command {
	case "map":
		c.Page++
		return nil
	case "mapb":
		if c.Page != 1 {
			c.Page--
		}

		return nil
	default:
		return nil
	}
}
