// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handler

import (
	"github.com/Leonargo404-code/find-my-brand/pkg/brand"
	"github.com/Leonargo404-code/find-my-brand/pkg/brand/service"
)

// Injectors from wire.go:

func Build() brand.Handlers {
	services := service.Must()
	handlers := Must(services)
	return handlers
}
