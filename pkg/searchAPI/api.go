package searchapi

import (
	"sync"

	"github.com/Leonargo404-code/find-my-brand/internal/env"
)

type (
	APIConfig struct {
		ApiKey string
		Engine string
	}

	SearchAPI interface {
		GoogleSearch(searchQuery, location string) ([]byte, error)
	}
)

var once sync.Once

func Must() (searchAPI SearchAPI) {
	once.Do(func() {
		searchAPI = &APIConfig{
			ApiKey: env.GetString(APIKEY),
			Engine: engineName,
		}
	})

	return searchAPI
}
