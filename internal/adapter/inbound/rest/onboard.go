package rest

import (
	"nftledger/internal/core/api"
	"nftledger/internal/core/model"

	"github.com/gofiber/fiber/v2"
)

type OnboardHandler struct {
	*api.Services
}

func NewOnboardHandler(
	api *api.Services,
) *OnboardHandler {
	return &OnboardHandler{
		api,
	}
}

func (h OnboardHandler) Save(c *fiber.Ctx) error {

	var onboard *model.Org = &model.Org{}
	err := c.BodyParser(onboard)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	txnHash, err := h.OnboardService.Onboard(onboard)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"txnHash": txnHash})
}
