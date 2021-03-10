package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "frete-rapido-desafio",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file: ", err)
	}
}
