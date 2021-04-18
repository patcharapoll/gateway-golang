package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gateway-golang/internal/graph/generated"
	"gateway-golang/internal/graph/model"
	"gateway-golang/internal/infrastructure/middleware"
	_model "gateway-golang/internal/model"
	"gateway-golang/internal/utils"
	service_v1 "gateway-golang/pkg/service/v1"
	"time"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) PingService(ctx context.Context, input *model.PingPong) (*bool, error) {
	_, err := r.grpcClient.PingPongServiceClient.StartPing(ctx, &service_v1.PingPong{
		Msg:  input.Msg,
		Ball: int32(input.Ball),
	})

	if err != nil {
		return nil, utils.ParseErrorResponse(err)
	}

	success := true
	return &success, nil
}

func (r *queryResolver) CustomerProfile(ctx context.Context, input model.DipChip) (*model.CustomerProfile, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StaffProfile(ctx context.Context) (*model.StaffProfile, error) {
	return &model.StaffProfile{
		Titile:       "MR",
		FirstName:    "FirstExample",
		LastName:     "LastExample",
		Team:         "RM2-1/IIP01",
		MobileNumber: "0801112222",
		IcLicenseNo:  "RM-12345",
		Email:        "rm001@gmail.com",
	}, nil
}

func (r *queryResolver) Login(ctx context.Context, input model.NewLogin) (*model.RmLogin, error) {
	encodedToken := middleware.ForAuthorizationContext(ctx)
	authClaimsToken, err := utils.DecodeToken(encodedToken)
	if err != nil {
		return nil, err
	}

	payload := authClaimsToken.ServicePayload

	context := context.Background()
	request := &service_v1.LoginRequest{
		UserId: payload.UserID,
		Pin:    input.Pin,
		// MobileNo: "1234",
		// Email:    "124",
		// DeviceInformation: &service_v1.DeviceInformation{
		// 	DeviceId:    "1",
		// 	DeviceModel: "2",
		// 	DeviceOs:    "3",
		// 	FcmToken:    "4",
		// },
	}

	res, err := r.grpcClient.LoginServiceClient.Login(context, request)

	if err != nil {
		return nil, err
	}

	token := utils.GenerateToken(_model.ServicePayload{}, time.Now().Add(r.config.TokenDuration).Unix())
	result := &model.RmLogin{
		Success: res.GetSuccess(),
		Token:   &token,
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
