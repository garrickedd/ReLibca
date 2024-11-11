package repository

type BaseRepository[T any] interface {
	Create(entity *T) error
	FindById(id string) (*T, error)
	// FindByName(name string) (*T, error)
	Update(entity *T) error
	Delete(id string) error
	GetAll() ([]T, error)
}
