package main

import (
	"os"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/shared/app"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "catalogs-api",
	Short:            "catalogs-api based on vertical slice architecture",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp().Run()
	},
}

// https://github.com/swaggo/swag#how-to-use-it-with-gin

// @contact.name Mehdi Hadeli
// @contact.url https://github.com/mehdihadeli
// @title Catalogs Api
// @version 1.0
// @description Catalogs Api.
func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
