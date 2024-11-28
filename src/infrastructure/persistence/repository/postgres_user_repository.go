package repository

import (
	"context"

	"github.com/garickedd/ReLibca/src/api/config"
	modeluser "github.com/garickedd/ReLibca/src/domain/model/user_aggregate"
	"github.com/garickedd/ReLibca/src/infrastructure/persistence/database"
	"github.com/garickedd/ReLibca/src/shared/constant"
	"github.com/garickedd/ReLibca/src/shared/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const userFilterExp string = "username = ?"
const countFilterExp string = "count(*) > 0"

type PostgresUserRepository struct {
	*BaseRepository[modeluser.User]
}

func NewUserRepository(cfg *config.Config) *PostgresUserRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return &PostgresUserRepository{BaseRepository: NewBaseRepository[modeluser.User](cfg, preloads)}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, u modeluser.User) (modeluser.User, error) {

	roleId, err := r.GetDefaultRole(ctx)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return u, err
	}
	tx := r.database.WithContext(ctx).Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return u, err
	}
	err = tx.Create(&modeluser.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return u, err
	}
	tx.Commit()
	return u, nil
}

func (r *PostgresUserRepository) FetchUserInfo(ctx context.Context, username string, password string) (modeluser.User, error) {
	var user modeluser.User
	err := r.database.WithContext(ctx).
		Model(&modeluser.User{}).
		Where(userFilterExp, username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *PostgresUserRepository) ExistsEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&modeluser.User{}).
		Select(countFilterExp).
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) ExistsUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&modeluser.User{}).
		Select(countFilterExp).
		Where(userFilterExp, username).
		Find(&exists).
		Error; err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&modeluser.User{}).
		Select(countFilterExp).
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) GetDefaultRole(ctx context.Context) (roleId int, err error) {

	if err = r.database.WithContext(ctx).Model(&modeluser.Role{}).
		Select("id").
		Where("name = ?", constant.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
