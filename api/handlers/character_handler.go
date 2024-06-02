package handlers

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/Danendz/genshin-api-go/types"
	"github.com/gofiber/fiber/v3"
)

type CharacterHandler struct {
	characterStore db.CharacterStore
}

func NewCharacterHandler(characterStore db.CharacterStore) *CharacterHandler {
	return &CharacterHandler{
		characterStore: characterStore,
	}
}

func (h *CharacterHandler) HandleGetCharacters(ctx fiber.Ctx) error {
	characters, err := h.characterStore.GetCharacters(ctx.Context())

	if err != nil {
		return err
	}

	if characters == nil {
		characters = make([]*types.Character, 0)
	}

	return ctx.JSON(api.NewApiResponse("characters fetched successfully", characters, true))
}

func (h *CharacterHandler) HandleGetCharacter(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	character, err := h.characterStore.GetCharacter(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("character fetched successfully", character, true))
}

func (h *CharacterHandler) HandleCreateCharacter(ctx fiber.Ctx) error {
	var character *types.CharacterCreateParams

	if err := ctx.Bind().Body(&character); err != nil {
		return err
	}

	if errors := validator.Validate(character); len(errors) > 0 {
		return ctx.JSON(api.NewApiResponse("invalid character", errors, false))
	}

	createdCharacter, err := h.characterStore.CreateCharacter(ctx.Context(), character)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("character created successfully", createdCharacter, true))
}

func (h *CharacterHandler) HandleDeleteCharacter(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	if err := h.characterStore.DeleteCharacter(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("character deleted successfully", nil, true))
}

func (h *CharacterHandler) HandleUpdateCharacter(ctx fiber.Ctx) error {
	var (
		id = ctx.Params("id")
		values *types.CharacterUpdateParams
	)

	if err := ctx.Bind().Body(&values); err != nil {
		return err
	}

	updatedCharacter, err := h.characterStore.UpdateCharacter(ctx.Context(), id, values)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("character updated successfully", updatedCharacter, true))
}