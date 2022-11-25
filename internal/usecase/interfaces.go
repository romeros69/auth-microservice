package usecase

import (
	"auth-microservice/internal/entity"
	"context"
)

type (
	UserRp interface {
		StoreUser(context.Context, entity.User) error
		GetUserByEmail(context.Context, string) (entity.User, error)
		CheckExistenceByEmail(context.Context, string) (bool, error)
	}

	UserContract interface {
		StoreUser(context.Context, entity.User) error
		GetUserByEmail(context.Context, string) (entity.User, error)
		CheckExistenceByEmail(context.Context, string) (bool, error)
	}

	JwtContract interface {
		CompareUserPassword(context.Context, entity.User) error
		GenerateToken(string) (string, error)
		CheckToken(string) (string, error)
	}
)
