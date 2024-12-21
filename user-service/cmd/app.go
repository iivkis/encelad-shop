package main

import (
	"encelad-shared/adapters/database"
	"encelad-shared/config"
	httphandler "enceland_user-service/internal/adapters/handler/http"
	"enceland_user-service/internal/adapters/repository"
	"enceland_user-service/internal/core/services"

	httpserver "encelad-shared/pkg/http_server"
)

func main() {
	db, err := database.NewPgxAdapter(config.Config().PostgresConnStr())
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserPostgresRepository(db)

	userService := services.NewUserService(userRepo)

	httpHandler := httphandler.NewHttpHandler(
		httphandler.NewUserHttpHandler(
			userService,
		),
	)

	httpServer := httpserver.NewHttpServer()
	if err := httpServer.Listen(
		config.Config().ServerAddress(),
		httpHandler,
	); err != nil {
		panic(err)
	}
}
