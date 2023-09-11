package service

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/gofrs/uuid"
)

type UserServiceImpl interface {
	Register(req *entity.UserRegister, ctx context.Context) (*entity.User, error)
	VerifyAccount(id uuid.UUID, ctx context.Context) (*entity.User, error)
	Login(req *entity.UserLogin, ctx context.Context) (*entity.ResponseLogin, error)
}
