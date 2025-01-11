package pokedex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PokeUrl = "https://pokeapi.co/api/v2/pokemon/%s/"
)

type Pokedex struct {
	Pokemons map[string]Pokemon
}
type Pokemon struct {
	Name            string `json:"name"`
	Id              int    `json:"id"`
	Base_experience int    `json:"base_experience"`
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokemons: make(map[string]Pokemon),
	}
}

func (p *Pokedex) Add(name string) (pokeRes Pokemon, err error) {

	pokemon := fmt.Sprintf(PokeUrl, name)

	res, err := http.Get(pokemon)

	if err != nil {
		fmt.Printf("\nerror in get!!: %v", err)
		return pokeRes, err
	}

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unknown pokemon!\n\nhttp status error: %d", res.StatusCode)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err = decoder.Decode(&pokeRes); err != nil {
		return pokeRes, err
	}

	return pokeRes, nil
}

func (p *Pokedex) RollCheck(roll int, stats *Pokemon) bool {

	switch {
	case roll == 1:
		fmt.Printf("\n\nCRITICAL FAILURE!\n\n")
		return false

	case roll == 20:
		fmt.Printf("\n\nCRITICAL SUCCESS!\n\n")
		return true

	case stats.Base_experience <= 64:
		return true
	case stats.Base_experience > 65 && stats.Base_experience <= 129:
		if roll < 6 {
			return false
		}
	case stats.Base_experience > 130 && stats.Base_experience <= 194:
		if roll < 11 {
			return false
		}
	case stats.Base_experience > 195:
		if roll < 16 {
			return false
		}
	}
	return true
}
