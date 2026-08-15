package fiber

import (
	"errors"

	"github.com/Novando/go-paket/util/response-constructor/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ResponseError(c *fiber.Ctx, err error) error {
	var stdErr dto.StandardResponse
	if errors.As(err, &stdErr) {
		return c.Status(stdErr.Status).JSON(dto.NewErrorResponse(stdErr))
	}

	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.NewErrorResponse(valErr))
	}
	return c.Status(fiber.StatusInternalServerError).JSON(dto.NewErrorResponse(err))
}
