package usecase

import (
	"context"

	"github.com/garickedd/ReLibca/src/api/config"
	"github.com/garickedd/ReLibca/src/domain/repository"
	"github.com/garickedd/ReLibca/src/shared/common"
	"github.com/garickedd/ReLibca/src/shared/logging"
)

type BaseUsecase[TEntity any, TCreate any, TUpdate any, TResponse any] struct {
	logger     logging.Logger
	repository repository.BaseRepository[TEntity]
}

func NewBaseUsecase[TEntity any, TCreate any, TUpdate any, TResponse any](cfg *config.Config, repository repository.BaseRepository[TEntity]) *BaseUsecase[TEntity, TCreate, TUpdate, TResponse] {
	logger := logging.NewLogger(cfg)
	return &BaseUsecase[TEntity, TCreate, TUpdate, TResponse]{
		repository: repository,
		logger:     logger,
	}
}

// Create
func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Create(ctx context.Context, req TCreate) (TResponse, error) {
	var response TResponse
	entity, _ := common.TypeConverter[TEntity](req)

	entity, err := u.repository.Create(ctx, entity)
	if err != nil {
		return response, err
	}

	response, _ = common.TypeConverter[TResponse](entity)
	return response, nil
}

// Update
func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Update(ctx context.Context, id int, req TUpdate) (TResponse, error) {
	var response TResponse
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)

	entity, err := u.repository.Update(ctx, id, updateMap)
	if err != nil {
		return response, err
	}
	response, _ = common.TypeConverter[TResponse](entity)

	return response, nil
}

// Read
func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) GetById(ctx context.Context, id int) (TResponse, error) {
	var response TResponse
	entity, err := u.repository.GetById(ctx, id)
	if err != nil {
		return response, err
	}
	return common.TypeConverter[TResponse](entity)
}

// Delete
func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Delete(ctx context.Context, id int) error {

	return u.repository.Delete(ctx, id)
}
