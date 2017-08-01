package main

import (
	"flag"
	"os"

	"github.com/alsx/enli-task/src/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// configMiddleware middleware adds a `DSN` to the context.
func configMiddleware(args map[string]interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for k, v := range args {
				c.Set(k, v)
			}
			return next(c)
		}
	}
}

func main() {
	const dsnHelp = `The Data Source Name (uri to connect to MySQL server)
	- example: username:password@protocol(address)/dbname?param=value
	- more: https://github.com/go-sql-driver/mysql#dsn-data-source-name
`
	dsn := flag.String("dsn", "", dsnHelp)
	flag.Parse()

	if *dsn == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// TODO: move to config
	const secret = "Some secret here"
	args := echo.Map{"dsn": *dsn, "secret": secret}
	e.Use(configMiddleware(args))

	e.Use(middleware.Static("dist"))
	e.File("/fb-signin/", "dist/index.html")
	e.File("/signin/", "dist/index.html")
	e.File("/signup/", "dist/index.html")

	// Unauthenticated route
	e.GET("/api/", handlers.VersionsHandler.List)
	e.GET("/api/v1/", handlers.IndexHandler.List)
	e.POST("/api/v1/signup/", handlers.UserHandler.SignUp)
	e.POST("/api/v1/signin/", handlers.UserHandler.SignIn)
	e.POST("/api/v1/fb-signup/", handlers.UserHandler.FacebookSignUp)

	// Restricted group
	r := e.Group("/api/v1/user/")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &handlers.UserClaims{},
		SigningKey: []byte(secret),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handlers.UserHandler.Info)

	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
