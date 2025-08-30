package core

import (
	"fmt"

	"github.com/ivnvMkhl/lekalo/config"
)

func prepareTemplateParams(templateParams []config.TemplateParam, userInputs map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})

	for _, param := range templateParams {
		userValue, exists := userInputs[param.Name]

		if exists {
			if param.Multiple {
				params[param.Name] = convertToJinja2Array(userValue)
			} else {
				params[param.Name] = convertToString(userValue)
			}
		} else {
			if param.Multiple {
				if param.Default != nil {
					params[param.Name] = convertToJinja2Array(param.Default)
				} else {
					params[param.Name] = []interface{}{}
				}
			} else {
				if param.Default != nil {
					params[param.Name] = convertToString(param.Default)
				} else {
					params[param.Name] = ""
				}
			}
		}
	}

	return params
}

func convertToJinja2Array(value interface{}) []interface{} {
	if value == nil {
		return []interface{}{}
	}

	switch v := value.(type) {
	case []string:
		result := make([]interface{}, len(v))
		for i, item := range v {
			result[i] = item
		}
		return result
	case []interface{}:
		return v
	case string:
		if v == "" {
			return []interface{}{}
		}
		return []interface{}{v}
	default:
		return []interface{}{fmt.Sprintf("%v", v)}
	}
}

func convertToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
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
