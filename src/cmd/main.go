// @title Relibca API
// @version 1.0
// @description RESTful API documentation for CoffeeBook Shop application.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/garrickedd/ReLibca/src/api/router"

	_ "github.com/garrickedd/ReLibca/src/docs"
	"github.com/garrickedd/ReLibca/src/infrastructure/database"
	"github.com/garrickedd/ReLibca/src/infrastructure/service"
	"github.com/spf13/viper"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("web.env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database, err := database.Postgres()
	if err != nil {
		log.Fatal(err)
	}
	router := router.NewRoute(database)
	server := service.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
