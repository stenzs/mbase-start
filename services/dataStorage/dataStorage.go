package dataStorage

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func GetHash(b []byte) string {
	hashing := sha256.New()
	hashing.Write(b)
	hash := base64.URLEncoding.EncodeToString(hashing.Sum(nil))
	return hash
}

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
	var b []byte
	var hash, path string

	b, err = SaveMultipartFileToBuffer(file)
	if err != nil {
		return err
	}

	hash = GetHash(b)
	path = fmt.Sprintf("./static/public/uploads/%s", hash)

	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = c.SaveFile(file, path)
		if err != nil {
			return err
		}
		return nil
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
