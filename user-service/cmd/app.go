package main

import (
	"encelad-shared/adapters/database"
	"encelad-shared/config"
	httphandler "enceland_user-service/internal/adapters/handler/http"
	"enceland_user-service/internal/adapters/repository"
	"enceland_user-service/internal/core/services"
	"fmt"
	"net/http"
)

func main() {
	db, err := database.NewPGXAdapter(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Config().PostgresDB.User,
		config.Config().PostgresDB.Password,
		config.Config().PostgresDB.Host,
		config.Config().PostgresDB.Port,
		config.Config().PostgresDB.Name,
	))
	if err != nil {
		panic(err)
	}

	UserRepository := repository.NewUserPostgresRepository(db)

	UserService := services.NewUserService(UserRepository)

	HTTPHandler := httphandler.NewHTTPHandler(
		httphandler.NewUserHTTPHandler(
			UserService,
		),
	)

	fmt.Println("server is runnig")
	if err := http.ListenAndServe(":80", HTTPHandler.Router); err != nil {
		panic(err)
	}

}
