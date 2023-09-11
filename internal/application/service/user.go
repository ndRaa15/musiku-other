package service

import (
	"context"
	"os"

	"github.com/Ndraaa15/musiku/global/email"
	"github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/global/jwt"
	"github.com/Ndraaa15/musiku/global/password"
	"github.com/Ndraaa15/musiku/global/validator"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/domain/service"
	"github.com/gofrs/uuid"
)

type UserService struct {
	Repository repository.UserRepositoryImpl
}

func NewUserService(ur repository.UserRepositoryImpl) service.UserServiceImpl {
	return &UserService{
		Repository: ur,
	}
}

func (us *UserService) Register(req *entity.UserRegister, ctx context.Context) (*entity.User, error) {
	err := validateRequestRegister(req)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	uuid := uuid.Must(uuid.NewV4())
	user := &entity.User{
		ID:       uuid,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Username: req.Username,
		Phone:    req.Phone,
	}

	user, err = us.Repository.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	mailer := email.NewMailClient()
	mailer.SetSubject("Email Verification")
	mailer.SetReciever(req.Email)
	mailer.SetSender(os.Getenv("CONFIG_SENDER_NAME"))
	mailer.SetBodyHTML(req.Username, os.Getenv("URL_VERIFY")+uuid.String())
	if err = mailer.SendMail(); err != nil {
		return nil, err
	}
	return user, nil
}

func validateRequestRegister(req *entity.UserRegister) error {
	if req.Name == "" {
		return errors.ErrNameRequired
	}

	if req.Email == "" || !validator.ValidateEmail(req.Email) {
		return errors.ErrInvalidEmail
	}

	if req.Password == "" || !validator.ValidatePassword(req.Password) {
		return errors.ErrInvalidPassword
	}

	if req.Username == "" {
		return errors.ErrUsernameRequired
	}

	if req.Phone == "" || !validator.ValidatePhone(req.Phone) {
		return errors.ErrInvalidPhoneNumber
	}
	return nil
}

func (us *UserService) VerifyAccount(id uuid.UUID, ctx context.Context) (*entity.User, error) {
	user, err := us.Repository.FindByID(id, ctx)
	if err != nil {
		return nil, err
	}

	if user.Status {
		return nil, errors.ErrAccountAlreadyVerified
	}

	user.Status = true
	user, err = us.Repository.Update(user, ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Login(req *entity.UserLogin, ctx context.Context) (*entity.ResponseLogin, error) {
	var res *entity.ResponseLogin

	user, err := us.Repository.FindByEmail(req.Email, ctx)
	if err != nil {
		return res, err
	}

	if !user.Status {
		return res, errors.ErrAccountNotVerified
	}

	if err := password.ComparePassword(req.Password, user.Password); err != nil {
		return res, errors.ErrInvalidPassword
	}

	jwt, err := jwt.EncodeToken(user)
	res = &entity.ResponseLogin{
		User:  user,
		Token: jwt,
	}
	return res, err
}
