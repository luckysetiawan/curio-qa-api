package routes

import (
	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/database"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/parser"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"github.com/luckysetiawan/curio-qa-api/pkg/usecase"
)

var mongoClient = database.NewMongoClient()
var redisClient = database.NewRedisClient()

var jsonPresenter = webserver.NewJsonPresenter()

var curioParser = parser.NewCurioParser()
var curioRepository = repository.NewCurioRepository(mongoClient, redisClient)
var curioUseCase = usecase.NewCurioUseCase(curioParser, jsonPresenter, curioRepository, userRepository)

var serverStatusRepository = repository.NewServerStatusRepository(mongoClient, redisClient)
var serverStatusUseCase = usecase.NewServerStatusUseCase(jsonPresenter, serverStatusRepository)

var userParser = parser.NewUserParser()
var userRepository = repository.NewUserRepository(mongoClient, redisClient)
var userUseCase = usecase.NewUserUseCase(userParser, jsonPresenter, userRepository)

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

	Get("/user/active", userUseCase.GetAllActiveUsers)
	Post("/user", userUseCase.Insert)
}

func init() {
	curioRoutes()
	serverStatusRoutes()
	userRoutes()
}
