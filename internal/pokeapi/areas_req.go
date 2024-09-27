package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListAreas(url *string) (Areas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if url != nil {
		fullURL = *url
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		var areas Areas
		err := json.Unmarshal(data, &areas)
		if err != nil {
			return Areas{}, err
		}

		return areas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Areas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Areas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Areas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Areas{}, err
	}

	var areas Areas
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return Areas{}, err
	}

	c.cache.Add(fullURL, data)

	return areas, nil
}

func (c *Client) GetArea(locationAreaName string) (Area, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		var area Area
		err := json.Unmarshal(data, &area)
		if err != nil {
			return Area{}, err
		}

		return area, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Area{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Area{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Area{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Area{}, err
	}

	var area Area
	err = json.Unmarshal(data, &area)
	if err != nil {
		return Area{}, err
	}

	c.cache.Add(fullURL, data)

	return area, nil
}
