package grpc

import (
	"auth-microservice/internal/controller/grpc/gen"
	"auth-microservice/internal/usecase"
	"context"
)

type AuthRegRpcServer struct {
	gen.UnimplementedAuthRegServer
	js usecase.JwtContract
	us usecase.UserContract
}

func NewAuthRegRpcServer(js usecase.JwtContract, us usecase.UserContract) *AuthRegRpcServer {
	return &AuthRegRpcServer{js: js, us: us}
}

func errCheckAuthResponse(msg string) *gen.CheckAuthResponse {
	return &gen.CheckAuthResponse{
		Result: false,
		Error:  msg,
	}
}

func errUserCredsResponse(msg string) *gen.UserCredentialsResponse {
	return &gen.UserCredentialsResponse{
		Email: "",
		Error: msg,
	}
}

func (a *AuthRegRpcServer) CheckAuthorization(ctx context.Context, req *gen.CheckAuthRequest) (*gen.CheckAuthResponse, error) {
	if req.Bearer == "" {
		return errCheckAuthResponse("token cannot be empty"), nil
	}
	email, err := a.js.CheckToken(req.Bearer)
	if err != nil {
		return errCheckAuthResponse(err.Error()), nil
	}
	isUser, err := a.us.CheckExistenceByEmail(ctx, email)
	if err != nil {
		return errCheckAuthResponse(err.Error()), nil
	}
	// TODO мб тут ошибка !
	if isUser {
		return &gen.CheckAuthResponse{
			Result: true,
			Error:  "",
		}, nil
	}
	return &gen.CheckAuthResponse{
		Result: false,
		Error:  "",
	}, nil
	//approved, err := a.us.CheckApproveUserByEmail(ctx, email)
	//if err != nil {
	//	return errCheckAuthResponse(err.Error()), nil
	//}
	//return &gen.CheckAuthResponse{
	//	Result: approved,
	//	Error:  "",
	//}, nil
}

func (a *AuthRegRpcServer) GetUserCredentials(ctx context.Context, req *gen.UserCredentialsRequest) (*gen.UserCredentialsResponse, error) {
	if req.Bearer == "" {
		return errUserCredsResponse("token cannot be empty"), nil
	}
	email, err := a.js.CheckToken(req.Bearer)
	if err != nil {
		return errUserCredsResponse(err.Error()), nil
	}
	//approved, err := a.us.CheckApproveUserByEmail(ctx, email)
	//if err != nil {
	//	return errUserCredsResponse(err.Error()), nil
	//}
	//if !approved {
	//	return errUserCredsResponse("user not approved"), nil
	//}
	user, err := a.us.GetUserByEmail(ctx, email)
	if err != nil {
		return errUserCredsResponse(err.Error()), nil
	}
	return &gen.UserCredentialsResponse{
		Email: user.Email,
		Error: "",
	}, nil
}
