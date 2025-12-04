package model

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Handler обрабатывает HTTP запросы
type Handler struct {
	config *Config
}

// NewHandler создает обработчик
func NewHandler(cfg *Config) *Handler {
	return &Handler{
		config: cfg,
	}
}

// ServeHTTP реализует интерфейс http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Только GET запросы
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Очищаем путь
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}

	// Полный путь к файлу
	fullPath := filepath.Join(h.config.GetStaticDir(), path)

	// Проверяем файл
	info, err := os.Stat(fullPath)
	if err != nil {
		h.serve404(w, path)
		return
	}

	// Создаем модель файла
	file := NewFileInfo(info.Name(), info.Size(), info.IsDir())

	// Отдаем файл
	if info.IsDir() {
		h.serveDir(w, fullPath)
	} else {
		h.serveFile(w, r, fullPath, file)
	}

	// Логируем
	log.Printf("[%s] %s %v", r.Method, path, time.Since(start))
}

// serveFile отдает файл
func (h *Handler) serveFile(w http.ResponseWriter, r *http.Request, path string, file *FileInfo) {
	w.Header().Set("Content-Type", file.GetContentType())
	http.ServeFile(w, r, path)
}

// serveDir показывает директорию
func (h *Handler) serveDir(w http.ResponseWriter, dirPath string) {
	entries, _ := os.ReadDir(dirPath)

	w.Header().Set("Content-Type", "text/html")

	html := "<h1>Directory listing</h1><ul>"
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			name += "/"
		}
		html += "<li><a href='" + name + "'>" + name + "</a></li>"
	}
	html += "</ul>"

	w.Write([]byte(html))
}

// serve404 показывает ошибку 404
func (h *Handler) serve404(w http.ResponseWriter, path string) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html")

	html := fmt.Sprintf("<h1>404 Not Found</h1><p>File %s not found</p>", path)
	w.Write([]byte(html))
}
