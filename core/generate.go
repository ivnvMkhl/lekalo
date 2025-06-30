package core

import (
	"fmt"
	"os"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/render"
)

func GenerateTemplate(templateName string, userInputs map[string]string) error {
	// Загружаем конфиги
	cfg, err := config.LoadConfigs()
	if err != nil {
		return err
	}

	template, exists := cfg.Templates[templateName]
	if !exists {
		return fmt.Errorf("шаблон '%s' не найден", templateName)
	}

	// Собираем все параметры (дефолтные + пользовательские)
	params := make(map[string]interface{})
	for _, param := range template.Params {
		value, ok := userInputs[param.Name]
		if !ok {
			value = param.Default // Используем дефолтное значение
		}
		params[param.Name] = value
	}

	// Создаём папки из `folders`
	if template.Folders != nil {
		for _, dirPath := range template.Folders {
			resolvedPath, err := render.ResolvePath(dirPath, params)
			if err != nil {
				return err
			}
			if err := render.EnsureDir(resolvedPath); err != nil {
				return err
			}
		}
	}

	// Генерируем файлы
	for _, file := range template.Files {
		// Рендерим путь
		filePath, err := render.ResolvePath(file.Path, params)
		if err != nil {
			return err
		}

		// Рендерим содержимое
		content, err := render.RenderString(file.Template, params)
		if err != nil {
			return err
		}

		// Создаём файл
		if err := render.EnsureDir(filePath); err != nil {
			return err
		}
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return err
		}

		fmt.Printf("Создан файл: %s\n", filePath)
	}

	return nil
}
