package handler

import (
	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
)

type handler struct {
	brandSvc brand.Services
}

func New(options ...Option) (brand.Handlers, error) {
	handler := &handler{}

	for _, option := range options {
		if err := option(handler); err != nil {
			return nil, err
		}
	}

	return handler, nil
}

func Must(brandSvc brand.Services) brand.Handlers {
	handler, err := New(WithService(brandSvc))
	if err != nil {
		panic(err)
	}
	return handler
}
