package pokedex

type Pokearea struct {
	Pokemons_on_area map[string]struct{}
}

func InitArea() (p *Pokearea) {
	return &Pokearea{
		Pokemons_on_area: make(map[string]struct{}),
	}
}

func (p *Pokearea) Refresh() {

	p.Pokemons_on_area = make(map[string]struct{})

}

func (p *Pokearea) Update(pokemon string) {

	p.Pokemons_on_area[pokemon] = struct{}{}

}
