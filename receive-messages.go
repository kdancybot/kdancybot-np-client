package main

import (
    "github.com/gorilla/websocket"
    "os"
    "encoding/json"
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

func handle_command(command string, url string) ([]byte, error) {
    if (command == "np") {
        return getDataFromGosumemory(url)
    }
    return nil, nil
}

func connection_handler(url string, gosumemory_url string, credentials []byte) (error) {
    // Establish a WebSocket connection
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    // Send the authentication payload to the server
    err = conn.WriteMessage(websocket.TextMessage, credentials)
    if err != nil {
        return err
    }

    // Handle incoming WebSocket messages
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            return err
        }

        // Process the received message
        response, err := handle_command(string(message), gosumemory_url)
        if err != nil {
            return err
        }
        
        if response != nil {
            err = conn.WriteMessage(websocket.TextMessage, response)
            if err != nil {
                return err
            }
        }
    }
}
