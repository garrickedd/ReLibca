package repository

import (
	"context"

	modelbook "github.com/garickedd/ReLibca/src/domain/model/book_aggregate"
	modelorder "github.com/garickedd/ReLibca/src/domain/model/order_aggregate"
	modelprd "github.com/garickedd/ReLibca/src/domain/model/product_aggregate"
	modeluser "github.com/garickedd/ReLibca/src/domain/model/user_aggregate"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	// GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}

type ProductRepository interface {
	BaseRepository[modelprd.Product]
}

type ProductTypeRepository interface {
	BaseRepository[modelprd.ProductType]
}

type ProductPropertyRepository interface {
	BaseRepository[modelprd.ProductProperty]
}

type BookRepository interface {
	BaseRepository[modelbook.Book]
}

type BookCategoryRepository interface {
	BaseRepository[modelbook.BookCategory]
}

type BookCategoryRelationRepository interface {
	BaseRepository[modelbook.BookCategoryRelation]
}

type OrderRepository interface {
	BaseRepository[modelorder.Order]
}

type OrderItemRepository interface {
	BaseRepository[modelorder.OrderItem]
}

type PromotionRepository interface {
	BaseRepository[modelorder.Promotion]
}

type UserRepository interface {
	ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	FetchUserInfo(ctx context.Context, username string, password string) (modeluser.User, error)
	GetDefaultRole(ctx context.Context) (roleId int, err error)
	CreateUser(ctx context.Context, u modeluser.User) (modeluser.User, error)
}

type RoleRepository interface {
	BaseRepository[modeluser.Role]
}

type UserRoleRepository interface {
	BaseRepository[modeluser.UserRole]
}
