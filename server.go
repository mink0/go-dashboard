package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/mink0/go-dashboard/dist/statik"
	"github.com/rakyll/statik/fs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init emv vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static(""))
	e.Use(middleware.CORS())
	
	// Init async jobs
	stats := statsJobRunner()

	// Route => handler
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(statikFS))))

	e.GET("/api/stats/cpu", func(c echo.Context) error {
		count, err := strconv.Atoi(c.QueryParam("count"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request param `count`")
		}

		return c.JSON(http.StatusOK, stats.cpu.Get(count))
	})

	// Start the server
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
