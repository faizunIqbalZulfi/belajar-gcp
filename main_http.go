package main

import (
	"belajar-gcp/domain/presenter"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func (s *Service) HTTPServeMain() {
	e := echo.New()
	// if config.Configuration.Env == config.EnvironmentDev {
	// 	e.Debug = true
	// 	e.Use(echoMiddleware.CORS())
	// }
	// e.Use(middleware.Metrics("go-core-lender", 0.3, 1.2, 5.0))
	// e.Use(echoMiddleware.Recover())

	// logMiddleware := dddmiddleware.Log{}
	// e.Use(logMiddleware.Logger)
	e.GET("/healthcheck", s.healthcheck)

	group := e.Group("/api/v1")

	pres := presenter.NewHTTPHandler(s.UseCase)
	pres.Mount(group)

	port := 8818
	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}

func (s *Service) healthcheck(c echo.Context) error {
	var err error
	if err = s.db.DB().Ping(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
