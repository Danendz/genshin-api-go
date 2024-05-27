package routes

import (
	"github.com/Danendz/genshin-api-go/api/handlers"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteParams struct {
	Client *mongo.Client
	DBcreds *db.DBCreds
}

func NewCharacterRoutes(router fiber.Router, params RouteParams) {
	characterHandler := handlers.NewCharacterHandler(
		db.NewMongoCharacterStore(params.Client, params.DBcreds),
	)

	router.Get("/", characterHandler.HandleGetCharacters)
	router.Post("/", characterHandler.HandleCreateCharacter)

	router.Get("/:id", characterHandler.HandleGetCharacter)
	router.Delete("/:id", characterHandler.HandleDeleteCharacter)
	router.Put("/:id", characterHandler.HandleUpdateCharacter)
}

func NewVisionRoutes(router fiber.Router, params RouteParams) {
	visionHandler := handlers.NewVisionHandler(db.NewMongoVisionStore(params.Client, params.DBcreds))

	router.Get("/", visionHandler.HandleGetVisions)
}