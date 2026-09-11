package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonFind(name string) (PokemonStruct, error) {
	if name == "" {
		return PokemonStruct{}, errors.New("Error: No parameter was given!")
	}
	url := baseURL + "/pokemon/" + name

	if data, exist := c.ClientCache.Get(url); exist {
		locationsResp := PokemonStruct{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return PokemonStruct{}, err
		}
		fmt.Print("Cache used Successfully :D\n")
		return locationsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonStruct{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("ListPokemon Error: %v", err)
		return PokemonStruct{}, err
	}
	if resp.StatusCode > 299 {
		return PokemonStruct{}, fmt.Errorf("Error: GET failed because %v,\n '%s' most likely wrong command", resp.Status, name)
	}
	//fmt.Print(resp.Status)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStruct{}, err
	}

	c.ClientCache.Add(url, dat)

	PokemonResp := PokemonStruct{}
	err = json.Unmarshal(dat, &PokemonResp)
	if err != nil {
		return PokemonStruct{}, err
	}

	return PokemonResp, nil

}

/*func (c Client) PokemonAdd(key string) {
	url := baseURL + "/pokemon/" + name


}*/
