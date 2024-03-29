package user_service

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/user"
	"amikom-pedia-api/module/user/user_repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository user_repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (userService *UserServiceImpl) Create(ctx context.Context, requestUser user.CreateRequestUser) user.ResponseUser {
	err := userService.Validate.Struct(requestUser)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	requestUserDomain := domain.User{
		Email:    requestUser.Email,
		Nim:      requestUser.Nim,
		Name:     requestUser.Name,
		Username: requestUser.Username,
		Bio:      requestUser.Bio,
		Password: requestUser.Password}

	result := userService.UserRepository.Create(ctx, tx, requestUserDomain)

	return helper.ToUserResponse(result)
}

func (userService *UserServiceImpl) Update(ctx context.Context) user.ResponseUser {
	//TODO implement me
	panic("implement me")
}

func (userService *UserServiceImpl) FindByUUID(ctx context.Context, uuid string) user.ResponseUser {
	//TODO implement me
	panic("implement me")
}

func (userService *UserServiceImpl) FindAll(ctx context.Context) []user.ResponseUser {
	//TODO implement me
	panic("implement me")
}
