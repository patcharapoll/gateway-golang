package resolver

import (
	"gateway-golang/internal/config"
	"gateway-golang/internal/graph/model"
	"gateway-golang/internal/infrastructure/grpcclient"

	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Module ...
var Module = fx.Provide(NewResolver)

// Resolver ...
type Resolver struct {
	config     *config.Configuration
	grpcclient *grpcclient.GRPC
	todos      []*model.Todo
}

// NewResolver ...
func NewResolver(
	config *config.Configuration,
	grpc *grpcclient.GRPC,
) *Resolver {
	resolver := &Resolver{
		config:     config,
		grpcclient: grpc,
	}
	return resolver
}
