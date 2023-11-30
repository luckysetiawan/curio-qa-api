package routes

import (
	"github.com/luckysetiawan/curio-qa-api/internal/database"
	"github.com/luckysetiawan/curio-qa-api/internal/webserver"
	"github.com/luckysetiawan/curio-qa-api/pkg/repository"
	"github.com/luckysetiawan/curio-qa-api/pkg/usecase"
)

var mongoClient = database.NewMongoClient()
var redisClient = database.NewRedisClient()

var jsonPresenter = webserver.NewJsonPresenter()

var serverStatusRepository = repository.NewServerStatusRepository(mongoClient, redisClient)
var serverStatusUseCase = usecase.NewServerStatusUseCase(jsonPresenter, serverStatusRepository)

func serverStatusRoutes() {
	Get("/status", serverStatusUseCase.GetStatus)
}

func init() {
	serverStatusRoutes()
}
