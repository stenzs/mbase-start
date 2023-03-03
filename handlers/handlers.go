package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func UpdateData(c *fiber.Ctx) error {

	file, err := c.FormFile("upload")
	airac := c.FormValue("airac")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = c.SaveFile(file, fmt.Sprintf("static/public/uploads/%s", file.Filename))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fmt.Println(airac)
	fmt.Println(file.Filename)

	return c.JSON(fiber.Map{
		"success": "create task",
	})

}
