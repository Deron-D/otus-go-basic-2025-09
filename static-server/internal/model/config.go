package model

// Config - настройки сервера
type Config struct {
	port      int    // приватное поле
	staticDir string // приватное поле
}

// NewConfig создает конфигурацию
func NewConfig() *Config {
	return &Config{
		port:      8080,
		staticDir: "./static",
	}
}

// GetPort возвращает порт
func (c *Config) GetPort() int {
	return c.port
}

// GetStaticDir возвращает директорию
func (c *Config) GetStaticDir() string {
	return c.staticDir
}

// SetPort устанавливает порт
func (c *Config) SetPort(port int) {
	c.port = port
}

// SetStaticDir устанавливает директорию
func (c *Config) SetStaticDir(dir string) {
	c.staticDir = dir
}
