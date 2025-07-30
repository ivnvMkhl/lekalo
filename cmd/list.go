package cmd

import (
	"fmt"

	"github.com/ivnvMkhl/lekalo/config"
	"github.com/ivnvMkhl/lekalo/locales"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: locales.T("list", "short"),
	Run: func(cmd *cobra.Command, args []string) {
		cfgs, err := config.LoadConfigs()
		if err != nil {
			fmt.Printf("%s: %s", locales.T("common", "error"), err)
			return
		}
		for name := range cfgs.Templates {
			fmt.Println("-", name)
		}
	},
}
