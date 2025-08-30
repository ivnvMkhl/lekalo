package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/locales"
	"github.com/ivnvMkhl/lekalo/render"
)

func GenerateTemplate(templateName string, userInputs map[string]interface{}) error {
	cfg, err := config.LoadConfigs()
	if err != nil {
		return err
	}

	template, exists := cfg.Templates[templateName]
	if !exists {
		return fmt.Errorf("%s '%s' %s", locales.T("common", "template"), templateName, locales.T("common", "not_found"))
	}

	params := prepareTemplateParams(template.Params, userInputs)

	if template.Folders != nil {
		for folderName, dirPath := range template.Folders {
			resolvedPath, err := render.ResolvePath(dirPath, params)
			if err != nil {
				return fmt.Errorf("%s %s: %w", locales.T("generate", "path_resolve_error"), folderName, err)
			}
			if err := os.MkdirAll(resolvedPath, 0755); err != nil {
				return fmt.Errorf("%s %s: %w", locales.T("generate", "folder_create_error"), resolvedPath, err)
			}
		}
	}

	for fileName, file := range template.Files {
		if err := processFile(file, params); err != nil {
			return fmt.Errorf("%s %s: %w", locales.T("generate", "file_processed_error"), fileName, err)
		}
	}

	return nil
}

func processFile(file config.FileTemplate, params map[string]interface{}) error {
	filePath, err := render.ResolvePath(file.Path, params)
	if err != nil {
		return fmt.Errorf("%s %s: %w", locales.T("generate", "path_resolve_error"), file.Path, err)
	}

	content, err := render.RenderString(file.Template, params)
	if err != nil {
		return fmt.Errorf("%s: %w", locales.T("generate", "render_error"), err)
	}

	switch file.Mode {
	case "always_overwrite":
		return overwriteFile(filePath, content)
	case "skip_if_exist":
		return handleSkipIfExist(filePath, content)
	case "append":
		return appendToFile(filePath, content)
	default:
		return handleInteractive(filePath, content)
	}
}

func handleSkipIfExist(filePath, content string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return writeFile(filePath, content)
	}
	fmt.Printf("%s: %s\n", locales.T("generate", "file_skip"), filePath)
	return nil
}

func handleInteractive(filePath, content string) error {
	if _, err := os.Stat(filePath); err == nil {
		input := UserInput(fmt.Sprintf("%s %s", filePath, locales.T("generate", "rewrite_message")))
		if UserInputToBool(input) {
			return overwriteFile(filePath, content)
		}
		return nil
	}
	return writeFile(filePath, content)
}

func writeFile(path string, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("%s %s: %w", locales.T("generate", "folder_create_error"), dir, err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("%s %s: %w", locales.T("generate", "file_write_error"), path, err)
	}
	fmt.Printf("%s: %s\n", locales.T("common", "file_created"), path)
	return nil
}

func overwriteFile(path string, content string) error {
	if err := writeFile(path, content); err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", locales.T("generate", "file_rewrited"), path)
	return nil
}

func appendToFile(path string, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("%s %s: %w", locales.T("generate", "folder_create_error"), dir, err)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("%s: %w", locales.T("generate", "file_open_failed"), err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("%s: %w", locales.T("generate", "file_append_failed"), err)
	}

	fmt.Printf("%s: %s\n", locales.T("generate", "file_appended"), path)
	return nil
}
