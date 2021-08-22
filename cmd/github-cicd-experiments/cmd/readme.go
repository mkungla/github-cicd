package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github-cicd-experiments/internal"
	"github-cicd-experiments/internal/readme"

	"github.com/spf13/cobra"
)

func Readme(app *internal.Application) *cobra.Command {
	var test, print bool
	cmd := &cobra.Command{}
	cmd.Use = "readme {update}"
	cmd.Short = "Manage repo README.md"
	cmd.Long = "Update readme file of this repository based on readme.yml and readme.tmpl"
	cmd.Args = cobra.MinimumNArgs(1)
	cmd.PersistentFlags().BoolVarP(&test, "test", "", false, "write readme output to README-TEST.md")
	cmd.PersistentFlags().BoolVarP(&print, "print", "", false, "print readme output to stdout")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "update":

			if os.Getenv("CI") != "true" && !test && !print {
				app.Log.Error("update can be only executed by github action. see update --help")
				return
			}
			cnf, err := app.Assets.ReadFile("assets/readme.yml")
			if err != nil {
				app.Log.Error(err)
				return
			}

			r := readme.New()
			if err := r.ReadConfigBytes(cnf); err != nil {
				app.Log.Error(err)
				return
			}
			r.AddLink(
				"git-repo",
				"https://github.com/mkungla/github-cicd-experiments",
				"GitHub CI/CD Experiments repository",
			)
			content, err := r.Render()
			if err != nil {
				app.Log.Fatal(err)
				return
			}

			if test {
				rpath := filepath.Join(app.WD, "README-TEST.md")
				f, err := os.OpenFile(rpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
				if err != nil {
					app.Log.Fatal(err)
					return
				}
				defer f.Close()

				if _, err := f.Write(content); err != nil {
					app.Log.Fatal(err)
					return
				}
				app.Log.Info("test README updated")
			} else if print {
				fmt.Println(string(content))
			} else {
				rpath := filepath.Join(app.WD, "README.md")
				f, err := os.OpenFile(rpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
				if err != nil {
					app.Log.Fatal(err)
					return
				}
				defer f.Close()

				if _, err := f.Write(content); err != nil {
					app.Log.Fatal(err)
					return
				}
				app.Log.Info("README updated")
			}

		default:
			log.Fatal("invalid arg")
		}
	}
	return cmd
}

type ReadmeFile struct {
}

// func newReadmeFile(app *internal.Application) *ReadmeFile {
// 	r := &ReadmeFile{}
// 	r.Path = filepath.Join(app.WD, "README.md")
