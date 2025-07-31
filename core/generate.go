package core

import (
	"fmt"
	"os"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/locales"
	"github.com/ivnvMkhl/lekalo/render"
)

func GenerateTemplate(templateName string, userInputs map[string]string) error {
	cfg, err := config.LoadConfigs()
	if err != nil {
		return err
	}

	template, exists := cfg.Templates[templateName]
	if !exists {
		return fmt.Errorf("%s '%s' %s", locales.T("common", "template"), templateName, locales.T("common", "not_found"))
	}

	params := make(map[string]interface{})
	for _, param := range template.Params {
		value, ok := userInputs[param.Name]
		if !ok {
			value = param.Default
		}
		params[param.Name] = value
	}

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

	for _, file := range template.Files {
		filePath, err := render.ResolvePath(file.Path, params)
		if err != nil {
			return err
		}

		content, err := render.RenderString(file.Template, params)
		if err != nil {
			return err
		}

		switch file.Mode {
		case "always_overwrite":
			{
				if err := overwriteFile(filePath, content); err != nil {
					return err
				}
			}
		case "skip_if_exist":
			{
				f, err := os.Open(filePath)
				if err != nil {
					if err := writeFile(filePath, content); err != nil {
						return err
					}
				} else {
					fmt.Printf("%s: %s\n", locales.T("generate", "file_skip"), filePath)
				}
				f.Close()
			}
		case "append":
			{
				if err := appendToFile(filePath, content); err != nil {
					return err
				}
			}
		default:
			{
				f, err := os.Open(filePath)
				if err == nil {
					input := UserInput(fmt.Sprintf("%s %s", filePath, locales.T("generate", "rewrite_message")))
					if confirm := UserInputToBool(input); confirm {
						if err := overwriteFile(filePath, content); err != nil {
							return err
						}
					}
				} else {
					if err := writeFile(filePath, content); err != nil {
						return err
					}
				}
				f.Close()
			}
		}

	}

	return nil
}

func writeFile(path string, content string) error {
	if err := render.EnsureDir(path); err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", locales.T("common", "file_created"), path)
	return nil
}

func overwriteFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", locales.T("generate", "file_rewrited"), path)
	return nil
}

func appendToFile(path string, content string) error {
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
