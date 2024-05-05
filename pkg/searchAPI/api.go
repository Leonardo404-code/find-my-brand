package searchapi

import (
	"os"
	"sync"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

type (
	APIConfig struct {
		ApiKey string
		Engine string
	}

	SearchAPI interface {
		GoogleSearch(searchQuery, location string, page int64) (*brand.Result, error)
	}
)

var once sync.Once

func Must() (searchAPI SearchAPI) {
	once.Do(func() {
		searchAPI = &APIConfig{
			ApiKey: os.Getenv(APIKEY),
			Engine: engineName,
		}
	})

	return searchAPI
}
