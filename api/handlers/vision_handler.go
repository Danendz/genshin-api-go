package handlers

import (
	"github.com/Danendz/genshin-api-go/api"
	"github.com/Danendz/genshin-api-go/db"
	"github.com/Danendz/genshin-api-go/types"
	"github.com/gofiber/fiber/v3"
)

type VisionHandler struct {
	visionStore db.VisionStore
}

func NewVisionHandler(visionStore db.VisionStore) *VisionHandler {
	return &VisionHandler{
		visionStore: visionStore,
	}
}

func (h *VisionHandler) HandleGetVisions(ctx fiber.Ctx) error {
	visions, err := h.visionStore.GetVisions(ctx.Context())

	if err != nil {
		return err
	}

	if visions == nil {
		visions = make([]*types.Vision, 0)
	}

	return ctx.JSON(api.NewApiResponse("visions fetched successfully", visions, true))
}

func (h *VisionHandler) HandleCreateVision(ctx fiber.Ctx) (err error) {
	var vision *types.VisionCreateParams

	if err = ctx.Bind().Body(&vision); err != nil {
		return err
	}

	if errors := validator.Validate(vision); len(errors) > 0 {
		return ctx.JSON(api.NewApiResponse("invalid vision", errors, false))
	}

	createdVision, err := h.visionStore.CreateVision(ctx.Context(), vision)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("vision created successfully", createdVision, true))
}

func (h *VisionHandler) HandleDeleteVision(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")

	if err = h.visionStore.DeleteVision(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("vision deleted successfully", nil, true))
}

func (h *VisionHandler) HandleUpdateVision(ctx fiber.Ctx) (err error) {
	var (
		id = ctx.Params("id")
		values *types.VisionUpdateParams
	)

	if err := ctx.Bind().Body(&values); err != nil {
		return err
	}

	if errors := validator.Validate(values); len(errors) > 0 {
		return ctx.JSON(api.NewApiResponse("invalid vision", errors, false))
	}

	updatedVision, err := h.visionStore.UpdateVision(ctx.Context(), id, values)

	if err != nil {
		return err
	}

	return ctx.JSON(api.NewApiResponse("vision updated successfully", updatedVision, true))
}