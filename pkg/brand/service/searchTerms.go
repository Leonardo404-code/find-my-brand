package service

import (
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

	result, err := s.SearchAPI.GoogleSearch(findBrandReq.Terms, location, 1)
	if err != nil {
		return nil, fmt.Errorf("failed in search with google: %v", err)
	}

	return result, nil
}
