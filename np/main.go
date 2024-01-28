package np

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
	"time"
)

func WsConnectionHandler() {
	config, err := LoadConfiguration("config.txt")
	if err != nil {
		log.Fatal("Error reading config file:", err)
		return
	}

	// Parse the server URL
	u, err := url.Parse(config.Host)
	if err != nil {
		log.Fatal("Error parsing server URL:", err)
		return
	}

	// Define your authentication payload (username and password)
	credentials, err := json.Marshal(config.Credentials)
	if err != nil {
		log.Fatal("Error marshalling credentials:", err)
		return
	}

	osu_urls := []string{config.StreamCompanionURL, config.GosumemoryURL}

	seconds_wait := 5 // This is fine until I implement good reconnection tactic
	for true {
		err = connection_handler(u.String(), osu_urls, credentials)
		log.Printf("Error happened during connection handling: %s", err)
		if strings.HasPrefix(err.Error(), "auth:") {
			log.Printf("Exiting...")
			return
		}
		log.Printf("Reconnecting in %d %s", seconds_wait, "seconds")
		time.Sleep(time.Duration(seconds_wait) * time.Second)
		// seconds_wait *= 2
	}
}
