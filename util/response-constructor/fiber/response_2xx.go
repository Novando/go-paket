package fiber

import (
	"strings"

	"github.com/Novando/go-paket/util/response-constructor/dto"
	"github.com/gofiber/fiber/v2"
)

func Response2xx(c *fiber.Ctx, res any, msg ...string) error {
	// Check if res is already a StandardResponse
	if sr, ok := res.(dto.StandardResponse); ok {
		return c.Status(sr.Status).JSON(sr)
	}

	if len(msg) < 1 {
		msg = append(msg, "Request Handled Successfully")
	}

	status := fiber.StatusOK
	code := "SUCCESS"
	if strings.Contains(strings.ToLower(msg[0]), "create") {
		status = fiber.StatusCreated
		code = "CREATED"
	}
	return c.Status(status).JSON(dto.StandardResponse{
		Code:    code,
		Message: msg[0],
		Value:   res,
		Status:  status,
	})
}

// NewArrayResponse creates a new StandardResponse with an array of values
func NewArrayResponse[T any, M dto.MetaNormal | dto.MetaCursor | any](val []T, meta M, msg ...string) dto.StandardResponse {
	if len(msg) < 1 {
		msg = append(msg, "Request Handled Successfully")
	}

	switch m := any(meta).(type) {
	case dto.MetaNormal:
		return dto.StandardResponse{
			Value:   dto.ArrayResponseWithMetaNormal[T]{Data: val, Meta: m},
			Code:    "SUCCESS",
			Message: msg[0],
		}
	case dto.MetaCursor:
		return dto.StandardResponse{
			Value:   dto.ArrayResponseWithMetaCursor[T]{Data: val, Meta: m},
			Code:    "SUCCESS",
			Message: msg[0],
		}
	default:
		return dto.StandardResponse{
			Value:   dto.ArrayResponse[T]{Data: val},
			Code:    "SUCCESS",
			Message: msg[0],
		}
	}
}
