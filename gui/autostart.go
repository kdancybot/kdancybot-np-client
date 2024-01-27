package gui

import (
	"log"
	"os"
	"path/filepath"
)

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
	// Get the path to the current executable
	path, _, err := shortcut.Read(GetShortcutPath())
	if err != nil {
		log.Println("Failed to find autostart entry:", err)
		return false
	}
	return true

	// _, err := os.Stat(GetShortcutPath())
	// return !os.IsNotExist(err)
}

// AddAutostart adds the application to autostart
func AddAutostart() {
	if err := shortcut.Make(GetRealPath(), GetShortcutPath(), ""); err != nil {
		log.Println("Failed to create autostart entry:", err)
		return
	}
	log.Println("Autostart entry added successfully!")
}

// RemoveAutostart removes the application from autostart
func RemoveAutostart() {
	// Remove the symlink to executable to disable autostart
	err := os.Remove(GetShortcutPath())
	if err != nil {
		log.Println("Error deleting old shortcut:", err)
		return
	}

	log.Println("Autostart entry removed successfully!")
}

// func main() {
//     if CheckAutostart() {
//         log.Println("Application is already added to autostart.")
//         // You can uncomment the line below to remove the application from autostart
//         // RemoveAutostart()
//     } else {
//         AddAutostart()
//     }
// }
