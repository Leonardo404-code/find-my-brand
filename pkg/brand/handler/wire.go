//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	"github.com/Leonargo404-code/find-my-brand/pkg/brand/service"
	"github.com/google/wire"
)

func Build() brand.Handlers {
	wire.Build(
		Must,
		service.Must,
	)

	return nil
}
