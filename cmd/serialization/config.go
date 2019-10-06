package main

import (
	cli "github.com/jawher/mow.cli"
)

var (
	appName = app.String(cli.StringOpt{
		Name:   "name",
		Desc:   "",
		EnvVar: "APP_NAME",
		Value:  "auth",
	})
	envName = app.String(cli.StringOpt{
		Name:   "env",
		Desc:   "",
		EnvVar: "APP_ENV",
		Value:  "local",
	})
	appLogLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "APP_LOG_LEVEL",
		Value:  "info",
	})

)
