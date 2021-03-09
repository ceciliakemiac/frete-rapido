package cmd

import (
	"fmt"
	// "os"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "frete-rapido",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	fmt.Println("Iniciando...")
}
