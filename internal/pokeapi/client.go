package pokeapi

import (
	"net/http";
	"encoding/json";
	"fmt";
	"io";
	"github.com/darkbreadmaker/Pokedex/internal/pokecache"
)
type Client struct {
	httpClient *http.Client
}
func NewClient() Client {
	return Client{
		httpClient: &http.Client{},
	}
}

type Location struct {
    Count    int    `json:"count"`
    Next     *string `json:"next"`
    Previous *string    `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

var cache pokecache.Cache = pokecache.NewCache(500)

func (c *Client) GetLocationArea(url string) Location {
	locationArea := Location{}
	value, ok := cache.Get(url)
	var body []byte
	if ok {
		body = value
	}else{
		res, err := http.Get(url)
		if err != nil {
			fmt.Errorf("Error making http connection: %s", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Errorf("Error reading data: %s", err)
		}	
	}	
	cache.Add(url, body)
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Errorf("Error unmarshalling data: %s", err)
	}
	return locationArea
} 
