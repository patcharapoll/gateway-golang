package container

import (
	"context"
	"gateway-golang/internal/config"
	"gateway-golang/internal/graph/resolver"
	"gateway-golang/internal/infrastructure/grpcclient"
	"gateway-golang/internal/infrastructure/http"

	"go.uber.org/fx"
)

// Container ...
type Container struct{}

// NewContainer ...
func NewContainer() *Container {
	return new(Container)
}

func (c *Container) configure() []fx.Option {
	return []fx.Option{
		config.Module,
		http.Module,
		grpcclient.Module,
		resolver.Module,
	}
}

func runApplication(
	lifecycle fx.Lifecycle,
	server *http.GinHTTPServer,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.Start(ctx)
				return nil
			},
			OnStop: func(ctx context.Context) error {

				return server.Stop(ctx)
			},
		},
	)
}

// Run ...
func (c *Container) Run() {
	option := append(c.configure(), fx.Invoke(runApplication))
	fx.New(option...).Run()
}
