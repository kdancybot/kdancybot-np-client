package gui

import (
	_ "embed"
	"os"

	"github.com/getlantern/systray"
)

// "os"

//go:embed icon.ico
var icon []byte

func Start() {
	SetUpApp()
	systray.Run(onReady, nil)
}

func onReady() {
	// icon, _ := os.ReadFile("icon.ico")
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("kdancybot")
	mAutostart := systray.AddMenuItem(getAutostartButtonTitle(), "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mAutostart.ClickedCh:
				addToAutostart()
				mAutostart.SetTitle(getAutostartButtonTitle())
			case <-mQuit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}

func getAutostartButtonTitle() string {
	if CheckAutostart() {
		return "Remove from autostart"
	} else {
		return "Add to autostart"
	}
}

func addToAutostart() {
	if CheckAutostart() {
		RemoveAutostart()
	} else {
		AddAutostart()
	}
}
