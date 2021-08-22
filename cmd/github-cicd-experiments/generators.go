// go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github-cicd-experiments/internal"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var app *internal.Application

func main() {
	var err error
	app, err = internal.App()
	if err != nil {
		log.Fatal(err)
	}
	if err := updateVersion(); err != nil {
		log.Fatal(err)
	}
}

// Versuin update.
var versionTmpl = template.Must(template.New("").Parse(`// Code generated by generators.go using 'go generate'. DO NOT EDIT.
// This file was generated at {{ .CreatedAt }}.
package main

const (
	// Version represents current version of this repo.
	Version = "{{ .Version }}"
	// Built represents date object when go gen was last execurted.
	Built = "{{ .CreatedAt }}"
)
`))

func updateVersion() error {
	// LATEST TAG
	tags, err := app.Repo.TagObjects()
	if err != nil {
		return err
	}
	var latest *object.Tag
	err = tags.ForEach(func(t *object.Tag) error {
		if latest == nil || t.Tagger.When.After(latest.Tagger.When) {
			latest = t
		}
		return nil
	})
	if err != nil {
		return err
	}

	ref, err := app.Repo.Head()
	if err != nil {
		return err
	}
	cIter, err := app.Repo.Log(&git.LogOptions{
		From:  ref.Hash(),
		Since: &latest.Tagger.When,
		Order: git.LogOrderCommitterTime,
	})

	var ccount int
	err = cIter.ForEach(func(c *object.Commit) error {
		ccount++
		return nil
	})

	var version string
	if ccount == 0 {
		version = latest.Name
	} else {
		version = fmt.Sprintf("%s+%d.%s", latest.Name, ccount, ref.Hash().String()[0:7])
	}

	v := struct {
		Filepath  string
		Package   string
		Version   string
		CreatedAt string
	}{
		Version:   version,
		CreatedAt: app.Now.Format(time.RFC3339),
	}
	f, err := os.Create(filepath.Join(app.WD, "cmd", "github-cicd-experiments", "version.go"))
	if err != nil {
		return err
	}
	return versionTmpl.Execute(f, v)
}
