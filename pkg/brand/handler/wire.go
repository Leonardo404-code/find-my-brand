//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	"github.com/Leonargo404-code/find-my-brand/pkg/brand/service"
	searchapi "github.com/Leonargo404-code/find-my-brand/pkg/searchAPI"
	"github.com/google/wire"
)

func Build() brand.Handlers {
	wire.Build(
		Must,
		service.Must,
		searchapi.Must,
	)

	return nil
}
