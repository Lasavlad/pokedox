package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error){
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil{
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)

	if ok {
		// cache hit
		fmt.Println("cache hit!")
		LocationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &LocationAreaResp)

		if err != nil {
			return LocationAreaResp, err
		}

		return LocationAreaResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	LocationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &LocationAreaResp)

	if err != nil {
		return LocationAreaResp, err
	}

	c.cache.Add(fullURL, dat)
	return LocationAreaResp, nil
}