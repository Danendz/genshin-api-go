package handlers

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/Danendz/genshin-api-go/types"
	"github.com/gofiber/fiber/v3"
)

type WeaponTypeHandler struct {
	weaponTypeStore db.WeaponTypeStore
}

func NewWeaponTypeHandler(weaponTypeStore db.WeaponTypeStore) *WeaponTypeHandler {
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
		weaponTypes = make([]*types.WeaponType, 0)
	}

	return ctx.JSON(api.NewApiResponse("weapon types fetched successfully", weaponTypes, true))
}

func (h *WeaponTypeHandler) HandleCreateWeaponType(ctx fiber.Ctx) (err error) {
	var weaponType *types.WeaponTypeCreateParams

	if err := ctx.Bind().Body(&weaponType); err != nil {
		return err
	}

	if errors := validator.Validate(weaponType); len(errors) > 0 {
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
		id = ctx.Params("id")
		values *types.WeaponTypeUpdateParams
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