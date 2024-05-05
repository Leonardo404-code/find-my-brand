package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

func (s *service) SearchTerms(
	findBrandReq *brand.FindBrandRequest,
	location string,
) (*brand.Result, error) {
	if strings.Trim(findBrandReq.Terms, " ") == "" {
		return nil, fmt.Errorf("no terms found in the request")
	}

	searchResult := new(brand.Result)

	result, err := s.SearchAPI.GoogleSearch(findBrandReq.Terms, location)
	if err != nil {
		return nil, fmt.Errorf("failed in search with google: %v", err)
	}

	if err := json.Unmarshal(result, searchResult); err != nil {
		return nil, fmt.Errorf("failed in unmarshal result: %v", err)
	}

	if len(searchResult.Ads) == 0 {
		return nil, fmt.Errorf("ads not found in this search")
	}

	return searchResult, nil
}
