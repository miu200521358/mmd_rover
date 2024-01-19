package main

import (
	"embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/miu200521358/mlib_go/pkg/front/mtheme"
	"github.com/miu200521358/mlib_go/pkg/utils/config"

	"github.com/miu200521358/mmd_rover/cmd/resources"
)

//go:embed resources/app_config.json
var appConfig embed.FS

func main() {
	appConfig := config.ReadAppConfig(appConfig)
	a := app.New()
	a.SetIcon(resources.AppIcon)
	a.Settings().SetTheme(&mtheme.MTheme{})
	w := a.NewWindow(fmt.Sprintf("%s %s", appConfig.AppName, appConfig.AppVersion))
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			widget.NewLabel("こんにちは、ファイン"),
			widget.NewLabel("これは日本語のラベルです"),
			widget.NewButton("これはボタンです", func() {
				dialog.ShowInformation("確認", "これはダイアログです", w)
			}),
			layout.NewSpacer(),
		),
	)

	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
