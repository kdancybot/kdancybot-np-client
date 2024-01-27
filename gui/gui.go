package gui

import (
	_ "embed"
	"os"

	"github.com/emersion/go-autostart"
	"github.com/getlantern/systray"
)

// "os"

//go:embed icon.ico
var icon []byte

var app *autostart.App

func Start() {
	app := &autostart.App{
		Name:        "kdancybot",
		DisplayName: "NPClient",
		Exec:        []string{GetRealPath()},
	}
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

// func startGosumemory() {
// 	log.Printf("started idk lol")
// 	// cmd := exec.Command("./gosumemory-no-window.exe")
// 	// if err := cmd.Run(); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }

func addToAutostart() {
	if CheckAutostart() {
		RemoveAutostart()
	} else {
		AddAutostart()
	}
	// cmd := exec.Command("./gosumemory-no-window.exe")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
