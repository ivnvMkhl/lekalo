package locales

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const DEFAULT_LANG = "en"

//go:embed */locales.json
var localeFS embed.FS

type Translations map[string]map[string]string

var translations Translations

func init() {
	lang := detectSystemLanguage()
	if err := LoadTranslations(lang); err != nil {
		_ = LoadTranslations(DEFAULT_LANG)
	}
}

func detectSystemLanguage() string {
	if lang := os.Getenv("LC_ALL"); lang != "" {
		return normalizeLang(lang)
	}
	if lang := os.Getenv("LANG"); lang != "" {
		return normalizeLang(lang)
	}
	return DEFAULT_LANG
}

func normalizeLang(lang string) string {
	lang = strings.Split(lang, ".")[0]
	lang = strings.Split(lang, "_")[0]

	return lang
}

func LoadTranslations(lang string) error {
	path := fmt.Sprintf("%s/locales.json", lang)
	file, err := localeFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load translations: %v", err)
	}

	err = json.Unmarshal(file, &translations)
	if err != nil {
		return fmt.Errorf("failed to parse translations: %v", err)
	}

	return nil
}

func T(group string, key string) string {
	if val, ok := translations[group][key]; ok {
		return val
	}
	return "Locale error " + group + "." + key
}
