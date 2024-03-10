// Package routes compile the necessary packages and handlers with the defined
// url to be added to the router.
package routes

import (
	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/database"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"github.com/luckysetiawan/curio-qa-api/pkg/usecase"
)

// Initialize database clients.
var mongoClient = database.NewMongoClient()
var redisClient = database.NewRedisClient()

// Initialize a JSON presenter.
var jsonPresenter = webserver.NewJsonPresenter()

// Initialize package handlers.
var curioParser = parser.NewCurioParser()
var curioRepository = repository.NewCurioRepository(mongoClient, redisClient)
var curioUseCase = usecase.NewCurioUseCase(curioParser, jsonPresenter, curioRepository, userRepository)

var serverStatusRepository = repository.NewServerStatusRepository(mongoClient, redisClient)
var serverStatusUseCase = usecase.NewServerStatusUseCase(jsonPresenter, serverStatusRepository)

var userParser = parser.NewUserParser()
var userRepository = repository.NewUserRepository(mongoClient, redisClient)
var userUseCase = usecase.NewUserUseCase(userParser, jsonPresenter, userRepository)

// Routes grouping according to the modules.
func curioRoutes() {
	Post("/curio/{receiverUsername}", curioUseCase.Insert, constant.RegisteredUser)
	Put("/curio/{curioID}", curioUseCase.UpdateStatus, constant.RegisteredUser)
}

func serverStatusRoutes() {
	Get("/server/status", serverStatusUseCase.GetStatus)
}

func userRoutes() {
	Post("/login", userUseCase.Login)
	Post("/logout", userUseCase.Logout)

	Get("/user/active", userUseCase.GetAllActiveUsers, constant.RegisteredUser)
	Get("/user", userUseCase.GetAll, constant.RegisteredUser)
	Get("/user/{username}", userUseCase.Find, constant.RegisteredUser)
	Post("/user", userUseCase.Insert)
}

// init calls all route groups when the server starts.
func init() {
	curioRoutes()
	serverStatusRoutes()
	userRoutes()
}
