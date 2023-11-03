package main

import (
    "log"
    "github.com/gorilla/websocket"
    "os"
    "encoding/json"
    "net/url"
)

type Config struct {
    Credentials struct {
        Client string `json:"client"`
        Password string `json:"password"`
    } `json:"credentials"`
    Host string `json:"host"`
    GosumemoryURL string `json:"gosumemory_url"`
}

func LoadConfiguration(file string) (Config, error) {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        return config, err
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config, err
}

func handle_command(command []byte, url string) ([]byte, error) {
    str_command := string(command)
    if (str_command == "np") {
        return getDataFromGosumemory(url)
    }
    return nil, nil
}

func main() {
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

    // Establish a WebSocket connection
    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("Error establishing WebSocket connection:", err)
        return
    }
    defer conn.Close()

    // Define your authentication payload (username and password)
    credentials, err := json.Marshal(config.Credentials)

    // Send the authentication payload to the server
    err = conn.WriteMessage(websocket.TextMessage, credentials)
    if err != nil {
        log.Fatal("Error sending authentication message:", err)
        return
    }

    // Handle incoming WebSocket messages
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading WebSocket message:", err)
            return
        }
        // Process the received message
        log.Printf("Received message: %s", message)

        response, err := handle_command(message, config.GosumemoryURL)
        if err != nil {
            log.Println("Error getting getting new data:", err)
            return
        }
        
        if response == nil {
            continue
        }
        err = conn.WriteMessage(websocket.TextMessage, response)
        if err != nil {
            log.Fatal("Error sending response:", err)
            return
        }
    }
}
