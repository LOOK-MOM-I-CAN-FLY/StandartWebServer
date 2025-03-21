package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!!!
	config *Config
	logger *logrus.Logger
	router *gin.Engine
}

// API constructor build API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

// Start http server/configure loggers, router, db conntection and etc
func (api *API) Start() error {
	//Trying to configure logger
	if err := api.configureLogger(); err != nil {
		return err
	}
	//Подтверждение того что логгер сконфигурирован
	api.logger.Info("Starting API server at port:", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configureRouter()
	//На этапе валидного завершения стартуем http.-сервер
	log.Fatal(api.router.Run(api.config.BindAddr))
	return nil
}
