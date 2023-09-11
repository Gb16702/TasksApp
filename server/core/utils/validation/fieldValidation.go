package validation

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"todoapp/core/utils"
)

func IsFieldEmpty(c *fiber.Ctx, field, err string) (bool, error) {
	if len(strings.TrimSpace(field)) == 0 {
		return true, utils.GenerateResponse(c, err, 400)
	}
	return false, nil
}
