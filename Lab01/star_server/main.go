package main

import (
	"fmt"
	"github.com/T1vz/itmo_clouds/Lab01/star_server/storage/postgres"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	pgDB := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPass, pgHost, pgPort, pgDB)
	fmt.Println(url)

	pgRepo := postgres.New(url)

	value := os.Getenv("APP_VALUE")

	dbInit := fmt.Sprintf("CREATE TABLE IF NOT EXISTS test (value varchar)")
	err := pgRepo.DoRequest(dbInit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	request := fmt.Sprintf("INSERT INTO test(value) VALUES (%s)", value)
	fmt.Println(request)
	err = pgRepo.DoRequest(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "3000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
