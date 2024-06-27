package routes

import (
	"github.com/Danendz/genshin-api-go/api"
	character2 "github.com/Danendz/genshin-api-go/api/handlers/character"
	dictionaries2 "github.com/Danendz/genshin-api-go/api/handlers/character/dictionaries"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/Danendz/genshin-api-go/db/character"
	"github.com/Danendz/genshin-api-go/db/character/dictionaries"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteParams struct {
	Client  *mongo.Client
	DBcreds *db.Creds
}

func NewCharacterRoutes(router fiber.Router, params RouteParams) {
	characterHandler := character2.NewCharacterHandler(
		character.NewMongoCharacterStore(params.Client, params.DBcreds),
	)

	router.Get("/", characterHandler.HandleGetCharacters)
	router.Get("/:id", characterHandler.HandleGetCharacter)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", characterHandler.HandleCreateCharacter)
	restricted.Delete("/:id", characterHandler.HandleDeleteCharacter)
	restricted.Put("/:id", characterHandler.HandleUpdateCharacter)
}

func NewVisionRoutes(router fiber.Router, params RouteParams) {
	visionHandler := dictionaries2.NewVisionHandler(dictionaries.NewMongoVisionStore(params.Client, params.DBcreds))

	router.Get("/", visionHandler.HandleGetVisions)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", visionHandler.HandleCreateVision)
	restricted.Delete("/:id", visionHandler.HandleDeleteVision)
	restricted.Put("/:id", visionHandler.HandleUpdateVision)
}

func NewWeaponTypeRoutes(router fiber.Router, params RouteParams) {
	weaponTypeHandler := dictionaries2.NewWeaponTypeHandler(dictionaries.NewMongoWeaponTypeStore(params.Client, params.DBcreds))

	router.Get("/", weaponTypeHandler.HandleGetWeaponTypes)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", weaponTypeHandler.HandleCreateWeaponType)
	restricted.Delete("/:id", weaponTypeHandler.HandleDeleteWeaponType)
	restricted.Put("/:id", weaponTypeHandler.HandleUpdateWeaponType)
}

func NewSkillTypeRoutes(router fiber.Router, params RouteParams) {
	weaponTypeHandler := dictionaries2.NewSkillTypeHandler(dictionaries.NewMongoSkillTypeStore(params.Client, params.DBcreds))

	router.Get("/", weaponTypeHandler.HandleGetSkillTypes)

	restricted := api.NewRestrictedApiRouter(router)
	restricted.Post("/", weaponTypeHandler.HandleCreateSkillType)
	restricted.Delete("/:id", weaponTypeHandler.HandleDeleteSkillType)
	restricted.Put("/:id", weaponTypeHandler.HandleUpdateSkillType)
}
