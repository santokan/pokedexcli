package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationRes, nil
}
