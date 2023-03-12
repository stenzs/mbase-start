package handlers

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"mbase/services/dataStorage"
	"mbase/services/validator"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

// customError returns custom 400 error
func customError(c *fiber.Ctx, err error, status int, description string) error {
	return c.Status(status).JSON(fiber.Map{
		"error":       err.Error(),
		"description": description,
	})
}

// successMessage returns success 200 message
func successMessage(c *fiber.Ctx, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "OK",
		"message": message,
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
	var files []*multipart.FileHeader
	var form *multipart.Form
	var airac string
	var b []byte

	airac = c.FormValue("airac")
	err = validator.ValidateAirac(airac)
	if err != nil {
		return customError(c, err, 400, "Невалидное значение airac")
	}

	form, err = c.MultipartForm()
	if err != nil {
		return customError(c, err, 400, "Не отправлена form-data")
	}

	files = form.File["upload"]
	if len(files) == 0 {
		err = errors.New("len of form.File[\"upload\"] < 0")
		return customError(c, err, 400, "Не загружено ни одного файла")
	}

	for _, file = range files {

		b, err = dataStorage.SaveMultipartFileToBuffer(file)
		if err != nil {
			return customError(c, err, 400, fmt.Sprintf("Ошибка сохранения файла %s в буфер",
				file.Filename))
		}

		err = validator.ValidateFile(b)
		if err != nil {
			return customError(c, err, 400, fmt.Sprintf("Ошибка валидации файла %s", file.Filename))
		}

		err = dataStorage.SaveFile(c, file)
		if err != nil {
			return customError(c, err, 400, fmt.Sprintf("Ошибка сохранения файла %s в хранилище",
				file.Filename))
		}

	}

	//hash = dataStorage.GetHash(b)
	//err = messageBroker.SendMessage(airac, hash)
	//if err != nil {
	//	return customError(c, err, 400, "Ошибка отправки сообщения в брокер")
	//}

	return successMessage(c, "Задача создана")

}
