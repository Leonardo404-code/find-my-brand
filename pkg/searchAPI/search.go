package searchapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

func (a *APIConfig) GoogleSearch(searchQuery, location string) (*brand.Result, error) {
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
		return nil, fmt.Errorf("failed in response")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed in read all")
	}

	result := new(brand.Result)

	if err := json.Unmarshal(body, result); err != nil {
		return nil, fmt.Errorf("failed in unmarshal result")
	}

	if len(result.Ads) == 0 {
		return nil, fmt.Errorf("ads not found in this search")
	}

	return result, nil
}
