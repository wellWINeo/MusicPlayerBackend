package main

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/handler"
	"log"
)

func main() {
	handler := handler.Handler{}
	srv := new(MusicPlayerBackend.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
