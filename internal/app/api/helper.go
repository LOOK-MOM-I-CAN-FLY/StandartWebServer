package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Пытаемся откунфигурировать наш API instane, а конкретнее поле Logger
func (a *API) configureLogger() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигугировать маршрутизатор (поле router API)
func (a *API) configureRouter() {
	a.router.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, This is my REST API")
	})
}
