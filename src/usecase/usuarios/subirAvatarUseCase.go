package usecase

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type SubirAvatarUseCase struct{}

func NewSubirAvatarUseCase() *SubirAvatarUseCase {
	return &SubirAvatarUseCase{}
}

func (uc *SubirAvatarUseCase) Execute(fileHeader *multipart.FileHeader) (string, error) {
	const maxSize = 5 * 1024 * 1024 // 5 MB

	if fileHeader.Size > int64(maxSize) {
		return "", fmt.Errorf("el archivo excede el tamaño máximo permitido de 2.5 MB")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return "", fmt.Errorf("tipo de archivo no permitido")
	}

	nombreArchivo := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	relPath := filepath.Join("uploads", "avatars", nombreArchivo)
	absPath := filepath.Join(".", relPath)

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return "", fmt.Errorf("error creando carpeta destino: %w", err)
	}

	if err := saveFile(fileHeader, absPath); err != nil {
		return "", fmt.Errorf("error guardando archivo: %w", err)
	}

	return "/" + relPath, nil
}

func saveFile(fileHeader *multipart.FileHeader, path string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err
}
