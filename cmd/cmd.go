package cmd

import (
	"fmt"
	"os"

	"github.com/ivnvMkhl/lekalo/locales"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "lekalo",
		Short: locales.T("cmd", "short"),
	}

	rootCmd.AddCommand(list)
	rootCmd.AddCommand(gen)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(locales.T("common", "error"), err)
		os.Exit(1)
	}
}
