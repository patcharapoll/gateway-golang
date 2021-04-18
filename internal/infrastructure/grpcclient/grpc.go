package grpcclient

import (
	"gateway-golang/internal/config"
	service_v1 "gateway-golang/pkg/service/v1"
	"log"
	"time"

	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Module ...
var Module = fx.Provide(NewGRPC)

// GRPC ...
type GRPC struct {
	config *config.Configuration

	// MY_SERVICE
	PingPongServiceClient service_v1.PingPongServiceClient
	LoginServiceClient    service_v1.LoginServiceClient
}

// NewGRPC ...
func NewGRPC(config *config.Configuration) (*GRPC, error) {
	g := &GRPC{
		config: config,
	}
	if err := g.connect(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *GRPC) connect() error {
	var (
		myService grpc.ClientConnInterface
		err       error
	)
	group := new(errgroup.Group)
	group.Go(func() error {
		myService, err = g.connectToService(g.config.EXAMPLEConnection)
		if err != nil {
			return err
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		return err
	}

	// myService, _ := g.connectToService(g.config.EXAMPLEConnection)
	g.PingPongServiceClient = service_v1.NewPingPongServiceClient(myService)
	g.LoginServiceClient = service_v1.NewLoginServiceClient(myService)

	return nil
}

func (g *GRPC) connectToService(host string) (grpc.ClientConnInterface, error) {
	conn, err := grpc.Dial(
		host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*5),
	)
	if err != nil {
		log.Fatalf("did not connect service: %v\n", err)
		return nil, err
	}
	log.Println("Connect to service on", host)
	return conn, nil
}
