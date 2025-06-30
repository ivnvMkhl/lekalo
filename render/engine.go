package render

import (
	"github.com/kluctl/go-jinja2"
)

func RenderString(template string, data map[string]interface{}) (string, error) {
	// Создаем движок Jinja2
	engine, err := jinja2.NewJinja2("render", 1) // 1 — количество worker'ов
	if err != nil {
		return "", err
	}
	defer engine.Close()

	// Рендерим шаблон, передавая данные через jinja2.WithGlobals()
	result, err := engine.RenderString(template, jinja2.WithGlobals(data))
	if err != nil {
		return "", err
	}

	return result, nil
}
