package main

import (
    "log"
    "encoding/json"
    "net/url"
    "time"
)

func main() {
    SelfUpdate()

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

    seconds_wait := 2
    for true {
        err = connection_handler(u.String(), config.GosumemoryURL, credentials)
        log.Printf("Error happened during connection handling: %s", err)
        log.Printf("Reconnecting in %d %s", seconds_wait, "seconds")
        time.Sleep(time.Duration(seconds_wait) * time.Second)
        seconds_wait *= 2
    }
}