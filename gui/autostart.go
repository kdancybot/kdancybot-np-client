package gui

import (
	"log"
	"os"
	"path/filepath"
)

func GetSymlinkPath() string {
	// Get user's home dir to enable autostart
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(os.Args[0])

	return homeDir + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\` + filename
}

// CheckAutostart checks if the application is already added to autostart
func CheckAutostart() bool {
	// Get the path to the current executable
	_, err := os.Stat(GetSymlinkPath())
	return !os.IsNotExist(err)
}

// AddAutostart adds the application to autostart
func AddAutostart() {
	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return
	}

	// Get through all symlinks to get real executable path
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		log.Println("Error getting real executable path:", err)
		return
	}

	os.Symlink(realPath, GetSymlinkPath())
	log.Println("Autostart entry added successfully!")
}

// RemoveAutostart removes the application from autostart
func RemoveAutostart() {
	// Remove the symlink to executable to disable autostart
	err := os.Remove(GetSymlinkPath())
	if err != nil {
		log.Println("Error deleting old symlink:", err)
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
