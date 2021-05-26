package main

import (
	"log"

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

	repos := repository.NewRepository()
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
