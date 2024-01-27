package gui

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emersion/go-autostart"
)

var app *autostart.App

func SetUpApp() {
	app = &autostart.App{
		Name:        "kdancybot",
                DisplayName: "NPClient",
                Exec:        []string{GetRealPath()},
        }
}

func GetShortcutPath() string {
	// Get user's home dir to enable autostart
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(os.Args[0])

	return filepath.Join(homeDir, `AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`, filename)
}

func GetRealPath() string {
	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return ""
	}
	log.Println(exePath)

	absolutePath, err := filepath.Abs(exePath)
	if err != nil {
		log.Println("Error getting executable path:", err)
		return ""
	}
	log.Println(absolutePath)

	// Get through all symlinks to get real executable path
	realPath, err := filepath.EvalSymlinks(absolutePath)
	if err != nil {
		log.Println("Error getting real executable path:", err)
		return ""
	}
	log.Println(realPath)

	return realPath
}

// CheckAutostart checks if the application is already added to autostart
func CheckAutostart() bool {
	return app.IsEnabled()
}

// AddAutostart adds the application to autostart
func AddAutostart() {
	if err := app.Enable(); err != nil {
		log.Fatal("Failed to add autostart entry:", err)
		return
	}
	log.Println("Autostart entry added successfully!")
}

// RemoveAutostart removes the application from autostart
func RemoveAutostart() {
	if err := app.Disable(); err != nil {
		log.Fatal("Failed to remove autostart entry:", err)
	}
	log.Println("Autostart entry removed successfully!")
}
