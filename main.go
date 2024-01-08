package main

import (
	"embed"

	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/miu200521358/mmd_rover/app"
	config "github.com/miu200521358/mmd_rover/app/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed wails.json
var wailsJSON string

func main() {
	productName := gjson.Get(wailsJSON, "Info.productName").String()
	productVersion := gjson.Get(wailsJSON, "Info.productVersion").String()

	// Create an instance of the app structure
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  productName + " " + productVersion,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.StartUp,
		Bind: []interface{}{
			app,
			&config.Config{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
