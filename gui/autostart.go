package gui

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

	// Get through all symlinks to get real executable path
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		log.Println("Error getting real executable path:", err)
		return ""
	}

	return realPath
}

func CreateShortcut(src, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", src)
	oleutil.CallMethod(idispatch, "Save")
	return nil
}

// CheckAutostart checks if the application is already added to autostart
func CheckAutostart() bool {
	// Get the path to the current executable
	_, err := os.Stat(GetShortcutPath())
	return !os.IsNotExist(err)
}

// AddAutostart adds the application to autostart
func AddAutostart() {
	CreateShortcut(GetRealPath(), GetShortcutPath())
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
