package dictionaries

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/api/handlers"
	"github.com/Danendz/genshin-api-go/db/character/dictionaries"
	dictionaries2 "github.com/Danendz/genshin-api-go/types/character/dictionaries"
	"github.com/gofiber/fiber/v3"
)

type SkillTypeHandler struct {
	skillTypeStore dictionaries.SkillTypeStore
}

func NewSkillTypeHandler(skillTypeStore dictionaries.SkillTypeStore) *SkillTypeHandler {
	return &SkillTypeHandler{
		skillTypeStore: skillTypeStore,
	}
}

func (h *SkillTypeHandler) HandleGetSkillTypes(ctx fiber.Ctx) (err error) {
	skillTypes, err := h.skillTypeStore.GetSkillTypes(ctx.Context())

	if err != nil {
		return err
	}

	if skillTypes == nil {
		skillTypes = make([]*dictionaries2.SkillType, 0)
	}

	return ctx.JSON(api.NewApiResponse("skill types fetched successfully", skillTypes, true))
}

func (h *SkillTypeHandler) HandleCreateSkillType(ctx fiber.Ctx) (err error) {
	var skillType *dictionaries2.SkillTypeCreateParams

	if err := ctx.Bind().Body(&skillType); err != nil {
		return err
	}

	if errors := handlers.Validator.Validate(skillType); len(errors) > 0 {
		return ctx.JSON(api.NewApiResponse("invalid skill type", errors, false))
	}

	createdSkillType, err := h.skillTypeStore.CreateSkillType(ctx.Context(), skillType)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("skill type created successfully", createdSkillType, true))
}

func (h *SkillTypeHandler) HandleDeleteSkillType(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")

	if err = h.skillTypeStore.DeleteSkillType(ctx.Context(), id); err != nil {
		return err

	}

	return ctx.JSON(api.NewApiResponse("skill type deleted successfully", nil, true))
}

func (h *SkillTypeHandler) HandleUpdateSkillType(ctx fiber.Ctx) (err error) {
	var (
		id     = ctx.Params("id")
		values *dictionaries2.SkillTypeUpdateParams
	)

	if err := ctx.Bind().Body(&values); err != nil {
		return err
	}

	updatedSkillType, err := h.skillTypeStore.UpdateSkillType(ctx.Context(), id, values)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("skill type updated successfully", updatedSkillType, true))
}
