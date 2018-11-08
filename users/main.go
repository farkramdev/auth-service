package main

import (
	"auth-service/users/common"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	// Calls startup logic
	common.StartUp()

	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.BodyLimit("2M"),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:9000",
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderContentType,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				echo.GET,
				echo.POST,
			},
			MaxAge: 3600,
		}),
	)

	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Register services
	// service.Auth(e.Group("/auth"))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
