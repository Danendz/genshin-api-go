package main

import (
	"context"
	"log"
	"os"

	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/api/routes"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(ctx fiber.Ctx, err error) error {
		return ctx.JSON(api.NewApiResponse(err.Error(), nil, false))
	},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//Mongo client
	dbcreds := db.NewDBCreds()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbcreds.DBURI).SetAuth(dbcreds.DBCREDS))

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	app := fiber.New(config)

	//Routes
	v1 := app.Group("/api/v1")

	routeParams := routes.RouteParams{
		Client: client,
		DBcreds: dbcreds,
	}

	routes.NewCharacterRoutes(v1.Group("/character"), routeParams)
	routes.NewVisionRoutes(v1.Group("/vision"), routeParams)
	routes.NewWeaponTypeRoutes(v1.Group("/weapon_type"), routeParams)

	app.Get("/health-check", func(ctx fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"message": "Ok",
		})
	})

	if err = app.Listen(port); err != nil {
		log.Fatal(err)
	}
}
