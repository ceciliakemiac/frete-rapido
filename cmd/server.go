package cmd

import (
	"log"
	"os"

	"github.com/ceciliakemiac/frete-rapido/api"
	"github.com/ceciliakemiac/frete-rapido/database"
	"github.com/spf13/cobra"
)

func startServer() {
	hasAccessExternalApi, err := api.HasAccessExternalApi()
	if err != nil || !hasAccessExternalApi {
		log.Fatal("Error Requesting Access to Frete Rapido External API: ", err)
	}

	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("startServer() Error Connecting to Database: ", err)
	}

	log.Println("Starting server...")

	server, err := api.NewServer(os.Getenv("API_ADDR"), db)
	if err != nil {
		log.Fatal("startServer() Error Creating New Server: ", err)
	}

	if err = server.Run(); err != nil {
		log.Fatal("startServer() Error Running Server: ", err)
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
