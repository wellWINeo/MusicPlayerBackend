package main

import (
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/handler"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Can't initialize config: %s", err.Error())
	}

	db, err := repository.NewMSSQLDB(repository.Config{
		Host: "localhost",
		Port: "1433",
		Username: "SA",
		Password: "A3F8CTvM4",
		DBName: "MusicPlayer",
	})

	if err != nil {
		log.Fatalf("Can't connect to DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(MusicPlayerBackend.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
