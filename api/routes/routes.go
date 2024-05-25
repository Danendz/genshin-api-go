package routes

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteParams struct {
	Client *mongo.Client
	DBcreds *db.DBCreds
}

func NewCharacterRoutes(router fiber.Router, params RouteParams) {
	characterHandler := api.NewCharacterHandler(
		db.NewMongoCharacterStore(params.Client, params.DBcreds),
	)

	router.Get("/", characterHandler.HandleGetCharacters)
	router.Post("/", characterHandler.HandleCreateCharacter)

	router.Get("/:id", characterHandler.HandleGetCharacter)
	router.Delete("/:id", characterHandler.HandleDeleteCharacter)
	router.Put("/:id", characterHandler.HandleUpdateCharacter)
}