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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Areas{}, err
	}

	var areas Areas
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return Areas{}, err
	}

	return areas, nil
}
