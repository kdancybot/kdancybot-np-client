package gui

import (
	_ "embed"
	"log"

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
	mAutostart := systray.AddMenuItem("Add to autostart", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mAutostart.ClickedCh:
				addToAutostart()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

// func startGosumemory() {
// 	log.Printf("started idk lol")
// 	// cmd := exec.Command("./gosumemory-no-window.exe")
// 	// if err := cmd.Run(); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }

func addToAutostart() {
	log.Printf("started idk lol")
	// cmd := exec.Command("./gosumemory-no-window.exe")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
