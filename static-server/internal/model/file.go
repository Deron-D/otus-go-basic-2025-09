package model

import (
	"fmt"
	"path/filepath"
)

// FileInfo - информация о файле
type FileInfo struct {
	Name        string
	Size        int64
	IsDir       bool
	contentType string // приватное поле
}

// NewFileInfo создает новую информацию о файле
func NewFileInfo(name string, size int64, isDir bool) *FileInfo {
	fi := &FileInfo{
		Name:  name,
		Size:  size,
		IsDir: isDir,
	}
	fi.contentType = fi.detectContentType()
	return fi
}

// GetContentType возвращает тип контента
func (fi *FileInfo) GetContentType() string {
	return fi.contentType
}

// GetSizeHuman возвращает размер в читаемом формате
func (fi *FileInfo) GetSizeHuman() string {
	if fi.Size < 1024 {
		return fmt.Sprintf("%d B", fi.Size)
	}
	if fi.Size < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(fi.Size)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(fi.Size)/(1024*1024))
}

// приватный метод
func (fi *FileInfo) detectContentType() string {
	ext := filepath.Ext(fi.Name)

	mime := map[string]string{
		".html": "text/html",
		".css":  "text/css",
		".js":   "application/javascript",
		".txt":  "text/plain",
		".json": "application/json",
	}

	if ct, ok := mime[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}
