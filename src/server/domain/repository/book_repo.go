package repository

import (
	"github.com/garrickedd/ReLibca/src/server/domain/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	BaseRepository[entity.Book]
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{db: db}
}

func (r *BookRepositoryImpl) Create(book *entity.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepositoryImpl) FindById(BookId string) (*entity.Book, error) {
	var book entity.Book
	if err := r.db.First(&book, "book_id = ?", BookId).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) Update(book *entity.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepositoryImpl) Delete(BookId string) error {
	return r.db.Delete(&entity.Book{}, "book_id = ?", BookId).Error
}

func (r *BookRepositoryImpl) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
