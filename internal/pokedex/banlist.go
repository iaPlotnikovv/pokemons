package pokedex

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Ban struct {
	mx       sync.Mutex
	Banlist  map[string]time.Time
	Cooldown time.Duration
}

func InitBan() *Ban {
	b := &Ban{
		Banlist:  make(map[string]time.Time),
		Cooldown: time.Duration(5 * time.Second),
	}
	go b.failLoop()
	return b
}

func (b *Ban) Fail(pokemon string) {
	b.mx.Lock()
	defer b.mx.Unlock()
	b.Banlist[pokemon] = time.Now()
}

func (b *Ban) failLoop() {

	ticker := time.NewTicker(b.Cooldown)

	for {

		<-ticker.C

		var PokemonsToCatch []string

		b.mx.Lock()
		for k, v := range b.Banlist {
			timePassed := time.Since(v)
			if timePassed > b.Cooldown {
				delete(b.Banlist, k)
				PokemonsToCatch = append(PokemonsToCatch, k)

			}

		}
		b.mx.Unlock()
		for _, p := range PokemonsToCatch {
			fmt.Fprintf(os.Stderr, "\n\n%s is available to catch!\n\n", p)
		}

	}
}
