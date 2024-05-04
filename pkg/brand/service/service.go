package service

import (
	"sync"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	searchapi "github.com/Leonargo404-code/find-my-brand/pkg/searchAPI"
)

type service struct {
	SearchAPI searchapi.SearchAPI
}

var once sync.Once

func Must(searchAPI searchapi.SearchAPI) (srvc brand.Services) {
	once.Do(func() {
		srvc = &service{
			SearchAPI: searchAPI,
		}
	})

	return srvc
}
