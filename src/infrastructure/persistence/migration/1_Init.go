package migration

import (
	"github.com/garickedd/ReLibca/src/api/config"
	modelbook "github.com/garickedd/ReLibca/src/domain/model/book_aggregate"
	modelorder "github.com/garickedd/ReLibca/src/domain/model/order_aggregate"
	modelprd "github.com/garickedd/ReLibca/src/domain/model/product_aggregate"
	modeluser "github.com/garickedd/ReLibca/src/domain/model/user_aggregate"
	"github.com/garickedd/ReLibca/src/infrastructure/persistence/database"
	"github.com/garickedd/ReLibca/src/shared/logging"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := database.GetDb()

	createTables(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// product
	product := modelprd.Product{}
	productType := modelprd.ProductType{}
	productProperty := modelprd.ProductProperty{}

	tables = addNewTable(database, productProperty, tables)
	tables = addNewTable(database, productType, tables)
	tables = addNewTable(database, product, tables)

	// book
	book := modelbook.Book{}
	bookCategory := modelbook.BookCategory{}
	bookCategoryRelation := modelbook.BookCategoryRelation{}

	tables = addNewTable(database, book, tables)
	tables = addNewTable(database, bookCategory, tables)
	tables = addNewTable(database, bookCategoryRelation, tables)

	// user
	user := modeluser.User{}
	role := modeluser.Role{}
	userRole := modeluser.UserRole{}

	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	// user
	promotion := modelorder.Promotion{}
	orderItem := modelorder.OrderItem{}
	order := modelorder.Order{}

	tables = addNewTable(database, promotion, tables)
	tables = addNewTable(database, orderItem, tables)
	tables = addNewTable(database, order, tables)

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
