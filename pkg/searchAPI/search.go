package searchapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

func (a *APIConfig) GoogleSearch(
	searchQuery,
	location string,
	page int64,
) (*brand.Result, error) {
	queryParams := url.Values{}
	queryParams.Add("api_key", a.ApiKey)
	queryParams.Add("engine", a.Engine)
	queryParams.Add("q", searchQuery)
	queryParams.Add("location", location)
	queryParams.Add("num", "100")
	queryParams.Add("page", fmt.Sprintf("%d", page))

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed in create a new request: %v", err)
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

	searchResult := new(brand.Result)

	if err := json.Unmarshal(body, searchResult); err != nil {
		return nil, fmt.Errorf("failed in unmarshal result: %v", err)
	}

	if len(searchResult.Ads) == 0 && page != 3 {
		page += 1
		a.GoogleSearch(searchQuery, location, page)
	}

	return searchResult, nil
}
