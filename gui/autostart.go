package gui

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// CheckAutostart checks if the application is already added to autostart
func CheckAutostart() bool {
	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return false
	}

	// Check the registry for the autostart entry
	key := `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
	cmd := exec.Command("reg", "query", key)
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error querying registry:", err)
		return false
	}

	// Check if the executable path exists in the registry output
	return strings.Contains(string(output), exePath)
}

// AddAutostart adds the application to autostart
func AddAutostart() {
	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		return
	}

	// Add registry key to enable autostart
	key := `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
	err = os.Setenv("KEY", key)
	if err != nil {
		log.Println("Error setting environment variable:", err)
		return
	}

	// Set the value to the path of the executable
	err = os.Setenv("VALUE", exePath)
	if err != nil {
		log.Println("Error setting environment variable:", err)
		return
	}

	// Run the PowerShell script to add the registry entry
	cmd := exec.Command("powershell", "-Command", "(Get-ItemEnv:$env:KEY).SetValue((Get-ItemEnv:$env:KEY).GetValueNames()[0], $env:VALUE)")
	if err := cmd.Run(); err != nil {
		log.Println("Error adding registry entry:", err)
		return
	}

	log.Println("Autostart entry added successfully!")
}

// RemoveAutostart removes the application from autostart
func RemoveAutostart() {
	// // Get the path to the current executable
	// exePath, err := os.Executable()
	// if err != nil {
	// 	log.Println("Error getting executable path:", err)
	// 	return
	// }

	// Remove the registry key to disable autostart
	key := `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
	err := os.Setenv("KEY", key)
	if err != nil {
		log.Println("Error setting environment variable:", err)
		return
	}

	// Run the PowerShell script to remove the registry entry
	cmd := exec.Command("powershell", "-Command", "(Get-ItemEnv:$env:KEY).DeleteValue((Get-ItemEnv:$env:KEY).GetValueNames()[0])")
	if err := cmd.Run(); err != nil {
		log.Println("Error removing registry entry:", err)
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
