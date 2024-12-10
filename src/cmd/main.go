package main

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/garrickedd/ReLibca/src/api/router"
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
