package handlers

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"mbase/services/dataStorage"
	"mbase/services/messageBroker"
	"mbase/services/validator"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

// customError returns custom 400 error
func customError(c *fiber.Ctx, err error, status int, description string) error {
	return c.Status(status).JSON(fiber.Map{
		"error":       true,
		"msg":         err.Error(),
		"description": description,
	})
}

// UpdateData godoc
// @Summary		Create new task
// @Tags			task
// @Accept			mpfd
// @Param			upload	formData	file	true	"uploaded file"
// @Param			airac	formData	int 	true	"airac"
// @Success		200		{string}	string	"answer"
// @Failure		400		{string}	string	"err"
// @Router /api/v1/task [post]
func UpdateData(c *fiber.Ctx) error {
	var err error
	var file *multipart.FileHeader
	var airac, hash string
	var b []byte

	airac = c.FormValue("airac")
	err = validator.ValidateAirac(airac)
	if err != nil {
		return customError(c, err, 400, "")
	}

	file, err = c.FormFile("upload")
	if err != nil {
		return customError(c, err, 400, "")
	}

	b, err = dataStorage.SaveMultipartFileToBuffer(file)
	if err != nil {
		return customError(c, err, 400, "")
	}

	err = validator.ValidateFile(b)
	if err != nil {
		return customError(c, err, 400, "")
	}

	err = dataStorage.SaveFile(c, file)
	if err != nil {
		return customError(c, err, 400, "")
	}

	hash = dataStorage.GetHash(b)
	err = messageBroker.SendMessage(airac, hash)
	if err != nil {
		return customError(c, err, 400, "")
	}

	return c.JSON(fiber.Map{
		"message": "task created",
	})

}
