package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	cache "github.com/iaPlotnikovv/pokemons/internal/pokecache"
)

func GetReq(url string, cache *cache.Cache) (apiResp APIResponse, err error) {

	//checking cache
	fmt.Printf("\nChecking cache for URL: %s\n", url)
	if data, ok := cache.Get(url); ok {

		if err := json.Unmarshal(data, &apiResp); err != nil {
			return apiResp, err
		}
		for _, loc := range apiResp.Results {
			fmt.Println(loc.Name)
		}

	} else {
		// request for data
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("\nerror in get!!: %v", err)
			return apiResp, err
		}
		if res.StatusCode != http.StatusOK {
			return APIResponse{}, fmt.Errorf("http status error: %d", res.StatusCode)
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)

		// storing data in cache

		if err = decoder.Decode(&apiResp); err != nil {
			return apiResp, err
		}

		cachedData, err := json.Marshal(apiResp)
		if err != nil {
			return apiResp, err
		}
		cache.Add(url, cachedData)

	}

	return apiResp, nil

}
