package rest

import (
	"nftledger/internal/core/api"
	"nftledger/internal/core/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler struct {
	*api.Services
}

func NewRequestHandler(
	api *api.Services,
) *RequestHandler {
	return &RequestHandler{
		api,
	}
}

func (h RequestHandler) Save(c *fiber.Ctx) error {

	var request *model.Request = &model.Request{}
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = h.RequestService.SaveRequest(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

func (h RequestHandler) Get(c *fiber.Ctx) error {

	id := c.Params("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var request = &model.Request{}
	request.ID = uint(ID)
	request, err = h.RequestService.GetRequest(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(request)
}

func (h RequestHandler) GetAll(c *fiber.Ctx) error {

	var request *model.Request = &model.Request{}
	if len(c.Body()) > 0 {
		err := c.BodyParser(request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
	}

	requests, err := h.RequestService.GetRequests(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(requests)
}

func (h RequestHandler) Update(c *fiber.Ctx) error {

	var request = map[string]model.Request{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
	}

	where, isExist := request["where"]
	if !isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "where is missing"})
	}

	update, isExist := request["update"]
	if !isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "updated data is missing"})
	}

	err = h.RequestService.UpdateRequest(&where, &update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("sucess")
}

func (h RequestHandler) Delete(c *fiber.Ctx) error {

	var request *model.Request
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = h.RequestService.DeleteRequest(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("success")
}
