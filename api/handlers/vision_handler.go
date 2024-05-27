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
