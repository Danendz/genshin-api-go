package api

import (
	"github.com/Danendz/genshin-api-go/api/routes"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var config = fiber.Config{
	ErrorHandler: func(ctx fiber.Ctx, err error) error {
		return ctx.JSON(NewApiResponse(err.Error(), nil, false))
	},
}

func NewFiberApp(client *mongo.Client, creds *db.Creds) {
	port := os.Getenv("PORT")
	app := fiber.New(config)

	//Routes
	v1 := app.Group("/api/v1")

	routeParams := routes.RouteParams{
		Client:  client,
		DBcreds: creds,
	}

	routes.NewCharacterRoutes(v1.Group("/character"), routeParams)
	routes.NewVisionRoutes(v1.Group("/vision"), routeParams)
	routes.NewWeaponTypeRoutes(v1.Group("/weapon_type"), routeParams)
	routes.NewSkillTypeRoutes(v1.Group("/skill_type"), routeParams)

	app.Get("/health-check", func(ctx fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"message": "Ok",
		})
	})

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
