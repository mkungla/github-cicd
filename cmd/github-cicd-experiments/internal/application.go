package internal

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Assets provided by addon

//go:embed assets
var fs embed.FS

type Application struct {
	started time.Time
	Now     time.Time
	WD      string // working directory
	Repo    *git.Repository
	Assets  embed.FS
	Log     *zap.SugaredLogger
	Version string
	Built   string
}

func App() (*Application, error) {
	var err error
	app := &Application{}
	app.started = time.Now()
	app.Assets = fs
	app.Now = time.Now().UTC()
	app.WD, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	app.Log = logger.Sugar()

	if _, err := os.Stat(filepath.Join(app.WD, ".git")); os.IsNotExist(err) {
		app.WD = filepath.Join(app.WD, "../../")
	}

	app.Repo, err = git.PlainOpen(app.WD)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *Application) Header() {
	fmt.Println(`
##########################################################################
# GitHub CI/CD Experiments
#
# This CLI application is providing commands to manage repo of 
# Github CI/CD Experiments: 
# (https://github.com/mkungla/github-cicd-experiments)
#
# Copyright © 2021 Marko Kungla 
# (https://github.com/mkungla). All rights reserved.
# 
# Version:      ` + a.Version + `
# Release date: ` + a.Built + `
##########################################################################
`)
}

func (a *Application) Sync() {
	if a.Log != nil {
		a.Log.Sync()
	}
	elapsed := time.Since(a.started)
	fmt.Println(`
##########################################################################
# elapsed ` + elapsed.String() + `
##########################################################################
`)
}
