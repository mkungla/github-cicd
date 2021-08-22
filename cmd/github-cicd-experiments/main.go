package main

//go:generate go run generators.go
import (
	"fmt"
	"github-cicd-experiments/cmd"
	"github-cicd-experiments/internal"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "github-cicd-experiments",
	Short:   "Github CI/CD Experiments managment cli",
	Run:     func(cmd *cobra.Command, args []string) {},
	Version: Version,
}

func main() {
	app, err := internal.App()
	app.Version = Version

	if err != nil {
		log.Fatal(err)
	}

	bd, err := time.Parse(time.RFC3339, Built)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	app.Built = bd.Format("2006-01-02")

	root.CompletionOptions.DisableDefaultCmd = true

	root.SetVersionTemplate(fmt.Sprintf("{{ .Version }} (%s)\n", app.Built))
	root.AddCommand(cmd.Readme(app))
	if err := root.Execute(); err != nil {
		app.Log.Error(err)
		os.Exit(1)
	}
}
