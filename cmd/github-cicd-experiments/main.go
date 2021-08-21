package main

//go:generate go run generators.go
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var app = &cobra.Command{
	Use:   "github-cicd-experiments",
	Short: "Github CI/CD Experiments managment cli",
	Long: `
This CLI application is providing commands to manage repo
of Github CI/CD Experiments:
(https://github.com/mkungla/github-cicd-experiments)`,
	Run:     func(cmd *cobra.Command, args []string) {},
	Version: Version,
}

func main() {
	bd, err := time.Parse(time.RFC3339, Built)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	app.SetVersionTemplate(fmt.Sprintf("{{ .Version }} (%s)\n", bd.Format("2006-01-02")))
	if err := app.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
