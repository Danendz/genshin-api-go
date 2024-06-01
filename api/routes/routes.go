package routes

import (
	"github.com/Danendz/genshin-api-go/api"
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
	router.Get("/:id", characterHandler.HandleGetCharacter)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", characterHandler.HandleCreateCharacter)
	restricted.Delete("/:id", characterHandler.HandleDeleteCharacter)
	restricted.Put("/:id", characterHandler.HandleUpdateCharacter)
}

func NewVisionRoutes(router fiber.Router, params RouteParams) {
	visionHandler := handlers.NewVisionHandler(db.NewMongoVisionStore(params.Client, params.DBcreds))

	router.Get("/", visionHandler.HandleGetVisions)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", visionHandler.HandleCreateVision)
	restricted.Delete("/:id", visionHandler.HandleDeleteVision)
	restricted.Put("/:id", visionHandler.HandleUpdateVision)
}

func NewWeaponTypeRoutes(router fiber.Router, params RouteParams) {
	weaponTypeHandler := handlers.NewWeaponTypeHandler(db.NewMongoWeaponTypeStore(params.Client, params.DBcreds))

	router.Get("/", weaponTypeHandler.HandleGetWeaponTypes)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", weaponTypeHandler.HandleCreateWeaponType)
	restricted.Delete("/:id", weaponTypeHandler.HandleDeleteWeaponType)
	restricted.Put("/:id", weaponTypeHandler.HandleUpdateWeaponType)
}