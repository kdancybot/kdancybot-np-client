package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/skratchdot/open-golang/open"
)

const version = "0.0.2"

// SelfUpdate updates the application (thx blackshark for this function <3)
func SelfUpdate() {
	log.Println("Checking Updates...")
	name, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "kdancybot/kdancybot-np-client")
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version)
		log.Println("Release notes:\n", latest.ReleaseNotes)
		full, _ := os.Executable()
		path, executable := filepath.Split(full)
		oldName := filepath.Join(path, "."+executable+".old")
		os.Remove(oldName)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release notes:\n", latest.ReleaseNotes)
		time.Sleep(3 * time.Second)
		open.Start(name)
		os.Exit(0)
	}
}