package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

func (s *service) SearchTerms(ctx context.Context, body []byte) (*brand.Result, error) {
	terms := new(brand.Terms)

	if err := json.Unmarshal(body, &terms); err != nil {
		return nil, fmt.Errorf("failed in unmarshal terms: %v", err)
	}

	if len(terms.Terms) == 0 {
		return nil, fmt.Errorf("no terms found in the request")
	}

	result, err := s.SearchAPI.GoogleSearch(terms.Terms[0], "Brazil")
	if err != nil {
		return nil, fmt.Errorf("failed in search with google: %v", err)
	}

	return result, nil
}
