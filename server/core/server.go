package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"todoapp/core/routes"
)

func HandleServerStart(port, DB_URL string) {

	app := fiber.New();
	app.Use(cors.New());

	routes.HandleAuthRoutes(app);
	app.Listen(port);

}
