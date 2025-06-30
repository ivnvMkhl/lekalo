package cmd

import (
	"fmt"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "Показать все шаблоны",
	Run: func(cmd *cobra.Command, args []string) {
		cfgs, err := config.LoadConfigs()
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		for name := range cfgs.Templates {
			fmt.Println("-", name)
		}
	},
}
