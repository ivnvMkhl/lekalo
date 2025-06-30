package render

import (
	"os"
	"path/filepath"
)

func EnsureDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0755)
}

func ResolvePath(templatePath string, data map[string]interface{}) (string, error) {
	return RenderString(templatePath, data) // Рендерим переменные в пути
}
