package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/core"
	"github.com/spf13/cobra"
)

var gen = &cobra.Command{
	Use:   "gen [template-name] [key=value...]",
	Short: "Сгенерировать файлы из шаблона",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		userInputs := make(map[string]string)

		// Парсим аргументы key=value
		for _, arg := range args[1:] {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) == 2 {
				userInputs[parts[0]] = parts[1]
			} else {
				fmt.Printf("Некорректный аргумент: %s. Ожидается формат key=value\n", arg)
				return
			}
		}

		// Загружаем конфиг
		cfg, err := config.LoadConfigs()
		if err != nil {
			fmt.Println("Ошибка загрузки конфига:", err)
			return
		}

		tpl, exists := cfg.Templates[templateName]
		if !exists {
			fmt.Printf("Шаблон '%s' не найден\n", templateName)
			return
		}

		// Запрашиваем недостающие параметры
		for _, param := range tpl.Params {
			if _, ok := userInputs[param.Name]; !ok {
				if param.Prompt != "" {
					fmt.Printf("%s: ", param.Prompt)
					var input string
					fmt.Scanln(&input)
					userInputs[param.Name] = input
				} else if param.Default != "" {
					userInputs[param.Name] = param.Default
				} else {
					fmt.Printf("Не указан обязательный параметр: %s\n", param.Name)
					return
				}
			}
		}

		if err := core.GenerateTemplate(templateName, userInputs); err != nil {
			fmt.Println("Ошибка генерации:", err)
			os.Exit(1)
		}
	},
}
