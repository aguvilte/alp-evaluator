package os

import (
	"mime/multipart"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"os"
	"io"
	"github.com/unlar/alp-evaluator/internal/core/ports"
)

type repository struct {
}

func NewRepo() ports.OSRepository {
	return &repository{}
}

func (r *repository) SaveFile(file multipart.File, header *multipart.FileHeader) error {
	newFile, err := os.Create(header.Filename)
	if err != nil {
		logger.Errorf("Hubo un error creando el archivo a guardar", err)
	}
	defer file.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		logger.Errorf("Hubo un error copiando el archivo a guardar", err)
	}

	return err
}
func (r *repository) DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		logger.Errorf("Hubo un error borrando el archivo", err)
	}

	return err
}
