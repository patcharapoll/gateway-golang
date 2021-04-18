package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway-golang/internal/graph/generated"
	"gateway-golang/internal/graph/model"
	_model "gateway-golang/internal/model"
	"gateway-golang/internal/utils"
	"math/rand"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{
			ID:   input.UserID,
			Name: "user " + input.UserID,
		},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) RmRegister(ctx context.Context, input *model.NewRegister) (*model.RmRegister, error) {
	result := true
	token := utils.GenerateToken(_model.ServicePayload{
		UserID: "914da998-9764-11eb-a8b3-0242ac130003",
	}, time.Now().Add(r.config.TokenDuration).Unix())

	return &model.RmRegister{
		Success: result,
		Token:   token,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
