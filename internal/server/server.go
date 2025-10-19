package server

import (
	"portfolioed/internal/config"
	"portfolioed/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config *config.Config
	router *fiber.App
}

func NewServer(config *config.Config) (*Server, error) {

	app := fiber.New(fiber.Config{})
	app.Use(middleware.RouteCacheMiddleware())

	// app.Use(cors.New(cors.Config{
	// 	AllowMethods: "GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS",
	// 	AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	// 	AllowOriginsFunc: func(origin string) bool {
	// 		u, err := url.Parse(origin)
	// 		if err != nil {
	// 			return false
	// 		}
	// 		h := u.Hostname()
	// 		return h == "localhost" || h == "127.0.0.1"
	// 	},
	// }))
	server := &Server{
		config: config,
		router: app,
	}
	server.setupStatics()
	return server, nil
}

func (server *Server) Start() error {
	return server.router.Listen(":" + server.config.AppPort)
}
func (server *Server) GetRouter() *fiber.App {
	return server.router
}
func (server *Server) GetConfig() *config.Config {
	return server.config
}
func (server *Server) setupStatics() {
	oneYearInSeconds := 31536000
	server.router.Static("/static", "./static", fiber.Static{
		MaxAge: oneYearInSeconds,
	})
}
