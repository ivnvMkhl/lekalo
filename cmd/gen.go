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
		userInputs := make(map[string]interface{})

		for _, arg := range args[1:] {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) == 2 {
				userInputs[parts[0]] = parts[1]
			} else {
				fmt.Printf("%s: %s. %s: key=value\n", locales.T("gen", "invalid_arg"), arg, locales.T("gen", "expected_format"))
				return
			}
		}

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

		for _, param := range tpl.Params {
			if _, ok := userInputs[param.Name]; !ok {
				if param.Prompt != "" {
					if param.Multiple {
						prompt := fmt.Sprintf("%s (%s): ", param.Prompt, locales.T("gen", "multiple_prompt_append"))
						input := core.UserInput(prompt)

						if input != "" {
							values := strings.Split(input, ",")
							trimmedValues := make([]string, len(values))
							for i, val := range values {
								trimmedValues[i] = strings.TrimSpace(val)
							}
							userInputs[param.Name] = trimmedValues
						} else if param.Default != nil {
							userInputs[param.Name] = convertDefaultToArray(param.Default)
						} else {
							userInputs[param.Name] = []string{}
						}
					} else {
						input := core.UserInput(param.Prompt + ": ")
						if input != "" {
							userInputs[param.Name] = input
						} else if param.Default != nil {
							userInputs[param.Name] = convertDefaultToString(param.Default)
						} else {
							userInputs[param.Name] = ""
						}
					}
				} else if param.Default != nil {
					if param.Multiple {
						userInputs[param.Name] = convertDefaultToArray(param.Default)
					} else {
						userInputs[param.Name] = convertDefaultToString(param.Default)
					}
				} else {
					if param.Multiple {
						userInputs[param.Name] = []string{}
					} else {
						userInputs[param.Name] = ""
					}
				}
			} else {
				if param.Multiple {
					if strValue, ok := userInputs[param.Name].(string); ok {
						if strValue != "" {
							values := strings.Split(strValue, ",")
							trimmedValues := make([]string, len(values))
							for i, val := range values {
								trimmedValues[i] = strings.TrimSpace(val)
							}
							userInputs[param.Name] = trimmedValues
						} else {
							userInputs[param.Name] = []string{}
						}
					}
				}
			}
		}

		if err := core.GenerateTemplate(templateName, userInputs); err != nil {
			fmt.Printf("%s: %s\n", locales.T("gen", "gen_error"), err)
			os.Exit(1)
		}
	},
}

func convertDefaultToArray(defaultValue interface{}) []string {
	if defaultValue == nil {
		return []string{}
	}

	switch v := defaultValue.(type) {
	case []string:
		return v
	case []interface{}:
		result := make([]string, len(v))
		for i, item := range v {
			if str, ok := item.(string); ok {
				result[i] = str
			} else {
				result[i] = fmt.Sprintf("%v", item)
			}
		}
		return result
	case string:
		if v == "" {
			return []string{}
		}
		return []string{v}
	default:
		return []string{fmt.Sprintf("%v", v)}
	}
}

func convertDefaultToString(defaultValue interface{}) string {
	if defaultValue == nil {
		return ""
	}

	switch v := defaultValue.(type) {
	case string:
		return v
	case []string:
		if len(v) > 0 {
			return v[0]
		}
		return ""
	case []interface{}:
		if len(v) > 0 {
			if str, ok := v[0].(string); ok {
				return str
			}
			return fmt.Sprintf("%v", v[0])
		}
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}
