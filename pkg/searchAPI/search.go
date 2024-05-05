package searchapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (a *APIConfig) GoogleSearch(searchQuery, location string) ([]byte, error) {
	queryParams := url.Values{}
	queryParams.Add("api_key", a.ApiKey)
	queryParams.Add("engine", a.Engine)
	queryParams.Add("q", searchQuery)
	queryParams.Add("location", location)

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed in create a new request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed in make request: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error in read body: %v", err)
	}

	return body, nil
}
