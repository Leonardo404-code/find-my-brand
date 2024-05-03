package handler

import "github.com/Leonargo404-code/find-my-brand/pkg/brand"

type Option func(*handler) error

func WithService(brandSvc brand.Services) Option {
	return func(h *handler) error {
		h.brandSvc = brandSvc
		return nil
	}
}
