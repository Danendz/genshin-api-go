package dictionaries

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/api/handlers"
	"github.com/Danendz/genshin-api-go/db/character/dictionaries"
	dictionaries2 "github.com/Danendz/genshin-api-go/types/character/dictionaries"
	"github.com/gofiber/fiber/v3"
)

type WeaponTypeHandler struct {
	weaponTypeStore dictionaries.WeaponTypeStore
}

func NewWeaponTypeHandler(weaponTypeStore dictionaries.WeaponTypeStore) *WeaponTypeHandler {
	return &WeaponTypeHandler{
		weaponTypeStore: weaponTypeStore,
	}
}

func (h *WeaponTypeHandler) HandleGetWeaponTypes(ctx fiber.Ctx) (err error) {
	weaponTypes, err := h.weaponTypeStore.GetWeaponTypes(ctx.Context())

	if err != nil {
		return err
	}

	if weaponTypes == nil {
		weaponTypes = make([]*dictionaries2.WeaponType, 0)
	}

	return ctx.JSON(api.NewApiResponse("weapon types fetched successfully", weaponTypes, true))
}

func (h *WeaponTypeHandler) HandleCreateWeaponType(ctx fiber.Ctx) (err error) {
	var weaponType *dictionaries2.WeaponTypeCreateParams

	if err := ctx.Bind().Body(&weaponType); err != nil {
		return err
	}

	if errors := handlers.Validator.Validate(weaponType); len(errors) > 0 {
		return ctx.JSON(api.NewApiResponse("invalid weapon type", errors, false))
	}

	createdWeaponType, err := h.weaponTypeStore.CreateWeaponType(ctx.Context(), weaponType)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("weapon type created successfully", createdWeaponType, true))
}

func (h *WeaponTypeHandler) HandleDeleteWeaponType(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")

	if err = h.weaponTypeStore.DeleteWeaponType(ctx.Context(), id); err != nil {
		return err

	}

	return ctx.JSON(api.NewApiResponse("weapon type deleted successfully", nil, true))
}

func (h *WeaponTypeHandler) HandleUpdateWeaponType(ctx fiber.Ctx) (err error) {
	var (
		id     = ctx.Params("id")
		values *dictionaries2.WeaponTypeUpdateParams
	)

	if err := ctx.Bind().Body(&values); err != nil {
		return err
	}

	updatedWeaponType, err := h.weaponTypeStore.UpdateWeaponType(ctx.Context(), id, values)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("weapon type updated successfully", updatedWeaponType, true))
}
