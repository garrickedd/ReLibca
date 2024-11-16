package migration

import (
	"github.com/garrickedd/ReLibca/src/server/config"
	"github.com/garrickedd/ReLibca/src/server/constant"
	"github.com/garrickedd/ReLibca/src/server/domain/model"
	"github.com/garrickedd/ReLibca/src/server/infrastructure/persistence/database"
	"github.com/garrickedd/ReLibca/src/server/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := database.GetDb()

	createTables(database)
	createDefaultInformation(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	user := model.User{}
	role := model.Role{}
	userRole := model.UserRole{}

	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := model.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := model.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := model.User{
		Username:     constant.DefaultUserName,
		FirstName:    "Test",
		LastName:     "Test",
		MobileNumber: "09111111111",
		Email:        "admin@admin.com"}
	pass := "123456789"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, r *model.Role) {
	exists := 0
	database.
		Model(&model.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *model.User, roleId int) {
	exists := 0
	database.
		Model(&model.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := model.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down_1() {}
