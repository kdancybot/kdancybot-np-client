package gui

import (
	_ "embed"
	"os"

	"github.com/getlantern/systray"
)

// "os"

//go:embed icon.ico
var icon []byte

// func handleRunningOsu() {
// 	var isRunningNow bool
// 	var cmds []exec.Cmd
// 	isRunningBefore := false
// 	for {
// 		isRunningNow = checkOsu()
// 		if isRunningBefore != isRunningNow {
// 			if isRunningNow {
// 				cmds = runNPClient()
// 			} else {
// 				stopNPClient(cmds)
// 			}
// 		}
// 		isRunningBefore = isRunningNow
// 	}
// }

func Start() {
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
