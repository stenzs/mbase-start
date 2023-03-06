package dataStorage

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"

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

// SaveMultipartFileToBuffer saves multipart file to buffer.
func SaveMultipartFileToBuffer(fh *multipart.FileHeader) (res []byte, err error) {
	var buffer bytes.Buffer
	var f multipart.File
	var ff *bufio.Writer

	ff = bufio.NewWriter(&buffer)
	f, err = fh.Open()

	defer func() {
		err := f.Close()
		if err == nil {
		}
	}()

	_, _ = copyZeroAlloc(ff, f)
	res = buffer.Bytes()
	return res, err
}

func copyZeroAlloc(w io.Writer, r io.Reader) (int64, error) {
	vBuf := copyBufPool.Get()
	buf := vBuf.([]byte)
	n, err := io.CopyBuffer(w, r, buf)
	copyBufPool.Put(vBuf)
	return n, err
}

var copyBufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}
