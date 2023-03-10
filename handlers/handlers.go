package handlers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"mbase/models"
	"mbase/services/dataStorage"
	"mbase/services/database"
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
		"error":       err.Error(),
		"description": description,
	})
}

// successMessage returns success 200 message
func successMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "OK",
		"message": message,
	})
}

// CreateTask godoc
// @Summary		Create new task
// @Tags			task
// @Accept			mpfd
// @Param			upload	formData	file	true	"uploaded file"
// @Param			airac	formData	int 	true	"airac"
// @Success		200		{string}	string	"answer"
// @Failure		400		{string}	string	"err"
// @Router /api/v1/task [post]
func CreateTask(c *fiber.Ctx) error {
	var err error
	var file *multipart.FileHeader
	var files []*multipart.FileHeader
	var form *multipart.Form
	var airac string
	var b []byte
	var filesHash []string
	var taskUuid = uuid.New()

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

		filesHash = append(filesHash, dataStorage.GetHash(b))

	}

	err = database.InsertTask(taskUuid)
	if err != nil {
		return customError(c, err, 400, "Ошибка записи в базу данных")
	}

	numAirac, _ := strconv.ParseInt(airac, 10, 16)
	message := models.Task{Uuid: taskUuid, PublishedAt: time.Now(), Airac: numAirac, Files: filesHash}

	err = messageBroker.SendMessage(message)
	if err != nil {
		return customError(c, err, 400, "Ошибка отправки сообщения брокеру")
	}

	return successMessage(c, 201, "Задача создана")

}

func UpdateTaskStatus(c *fiber.Ctx) error {

	payload := struct {
		TaskUuid uuid.UUID `json:"uuid"`
		Status   string    `json:"status"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return customError(c, err, 400, "Некорректные данные")
	}

	if err := database.UpdateTaskByUuid(payload.TaskUuid, payload.Status); err != nil {
		return customError(c, err, 400, "Ошибка подключения к бд")
	}

	return successMessage(c, 200, "Статус обновлен")

}
