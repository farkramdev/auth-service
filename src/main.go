package main

import (
	"crypto/tls"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"net"
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

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"prefix1.mongodb.net:27017",
			"prefix2.mongodb.net:27017",
			"prefix3.mongodb.net:27017"},
		Database: "authDatabaseName",
		Username: "user",
		Password: "pass",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	e := echo.New()
	e.Use(middleware.CORS())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", read)
	e.PUT("/users/:id", update)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
