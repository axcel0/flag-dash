package server

import (
	_ "github.com/blastertwist/flag-dash/docs"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /api
func (s *Server) InitializeSwagger(){
	s.fiber.Get("/swagger/*", swagger.HandlerDefault)
}