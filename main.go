package main

import "github.com/HungTP-Play/go-temp/app"

func main() {
	appConfigs := app.MergeAppConfigsWithEnv()
	tempApp := app.NewApp(appConfigs)

	tempApp.Start()
}
