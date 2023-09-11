package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) repository.UserRepositoryImpl {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *entity.User, ctx context.Context) (*entity.User, error) {
	if err := ur.db.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return nil, gorm.ErrRegistered
	}
	return user, nil
}

func (ur *UserRepository) FindByID(id uuid.UUID, ctx context.Context) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Debug().WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (ur *UserRepository) Update(user *entity.User, ctx context.Context, id uuid.UUID) (*entity.User, error) {
	if err := ur.db.Debug().WithContext(ctx).Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (ur *UserRepository) FindByEmail(email string, ctx context.Context) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}
