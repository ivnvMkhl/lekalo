package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/core"
	"github.com/ivnvMkhl/lekalo/locales"
	"github.com/spf13/cobra"
)

var gen = &cobra.Command{
	Use:   "gen [template-name] [key=value...]",
	Short: locales.T("gen", "short"),
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
				fmt.Printf("%s: %s. %s: key=value\n", locales.T("gen", "invalid_arg"), arg, locales.T("gen", "expected_format"))
				return
			}
		}

		// Загружаем конфиг
		cfg, err := config.LoadConfigs()
		if err != nil {
			fmt.Printf("%s: %s", locales.T("gen", "config_load_error"), err)
			return
		}

		tpl, exists := cfg.Templates[templateName]
		if !exists {
			fmt.Printf("%s '%s' %s\n", locales.T("common", "template"), templateName, locales.T("common", "not_found"))
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
					fmt.Printf("%s: %s\n", locales.T("gen", "required_param_error"), param.Name)
					return
				}
			}
		}

		if err := core.GenerateTemplate(templateName, userInputs); err != nil {
			fmt.Printf("%s: %s", locales.T("gen", "gen_error"), err)
			os.Exit(1)
		}
	},
}
