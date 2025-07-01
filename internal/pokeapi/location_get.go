package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocationsPokemon -
func (c *Client) GetSpecificLocation(locationURL string) (Location, error) {
	url := baseURL + "/location-area/" + locationURL

	if dat, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		if err := json.Unmarshal(dat, &locationsResp); err != nil {
			return Location{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	if err = json.Unmarshal(dat, &locationsResp); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
