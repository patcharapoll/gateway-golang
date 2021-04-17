package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway-golang/internal/graph/generated"
	"gateway-golang/internal/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PingService(ctx context.Context, input *model.PingPong) (*bool, error) {
	success := true
	return &success, nil
}

func (r *queryResolver) CustomerProfile(ctx context.Context, input model.DipChip) (*model.CustomerProfile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StaffProfile(ctx context.Context) (*model.StaffProfile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context, input model.NewLogin) (*model.RmLogin, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
