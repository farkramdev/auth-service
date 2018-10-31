package main

import (
	// "crypto/tls"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "gopkg.in/mgo.v2"
	// "net"
	"net/http"
	"strconv"
)

// User gggg
type User struct {
	ID        int    `json:"id" xml:"id"`
	Firstname string `json:"firstname" xml:"firstname"`
	Lastname  string `json:"lastname" xml:"lastname"`
}

var (
	users = map[int]*User{}
	seq   = 1
)

func createUser(c echo.Context) error {
	u := &User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func read(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func update(c echo.Context) error {

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))

	users[id].Firstname = u.Firstname

	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {

	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.BodyLimit("2M"),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:8080",
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

	// Register services
	service.Auth(e.Group("/auth"))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
