package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "lekalo",
		Short: "Генератор файлов по шаблонам",
	}

	rootCmd.AddCommand(list)
	rootCmd.AddCommand(gen)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
}
