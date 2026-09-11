package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(name string) (RespDeepLocation, error) {
	if name == "" {
		return RespDeepLocation{}, errors.New("Error: No parameter was given!")
	}
	url := baseURL + "/location-area/" + name

	if data, exist := c.ClientCache.Get(url); exist {
		locationsResp := RespDeepLocation{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		fmt.Print("Cache used Successfully :D\n")
		return locationsResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("ListPokemon Error: %v", err)
		return RespDeepLocation{}, err
	}
	if resp.StatusCode > 299 {
		return RespDeepLocation{}, fmt.Errorf("Error: GET failed because %v,\n '%s' most likely wrong command", resp.Status, name)
	}
	//fmt.Print(resp.Status)
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	c.ClientCache.Add(url, dat)

	DeepLocationsResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &DeepLocationsResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	return DeepLocationsResp, nil

}
