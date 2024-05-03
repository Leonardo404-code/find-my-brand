package service

import (
	"sync"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

type service struct{}

var once sync.Once

func Must() brand.Services {
	return new(service)
}
