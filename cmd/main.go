package main

import (
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/handler"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Can't initialize config: %s", err.Error())
	}

	if err := godotenv.Load(); err  != nil {
		logrus.Fatalf("Can't read .env file: %s", err.Error())
	}

	db, err := repository.NewMSSQLDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("Can't connect to DB: %s", err.Error())
	}

	mailConfig := service.AuthConfig{
		Host: viper.GetString("mail.host"),
		Port: viper.GetInt("mail.port"),
		MailBox: os.Getenv("MAIL_BOX"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Salt: os.Getenv("SALT"),
		TokenSecret: os.Getenv("SECRET"),
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos, mailConfig)
	handler := handler.NewHandler(services)
	srv := new(MusicPlayerBackend.Server)
	logrus.Printf("Running on port: %d", viper.GetInt("server.port"))
	if err := srv.Run(viper.GetInt("server.port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
