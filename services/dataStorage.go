package dataStorage

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

func MakeUploadFolder() {
	var err error

	if _, err = os.Stat("./static/public/uploads"); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir("./static/public/uploads", os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func SaveFile(c *fiber.Ctx, file *multipart.FileHeader) error {
	var err error

	err = c.SaveFile(file, fmt.Sprintf("./static/public/uploads/%s", file.Filename))
	if err != nil {
		return err
	}
	return nil
}
