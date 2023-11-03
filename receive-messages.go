package main

import (
    "log"
    "github.com/gorilla/websocket"
    "io/ioutil"
    "net/url"
)

func handle_command(command []byte) ([]byte, error) {
    str_command := string(command)
    if (str_command == "np") {
        return getDataFromGosumemory()
    }
    return nil, nil
}

func main() {
    // Specify the WebSocket server URL
    serverURL := "ws://localhost:1727/ws"

    // Parse the server URL
    u, err := url.Parse(serverURL)
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
    credentials, err := ioutil.ReadFile("config.txt")

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

        response, err := handle_command(message)
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
