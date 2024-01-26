package np

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

type Config struct {
	Credentials struct {
		Client   string `json:"client"`
		Password string `json:"password"`
	} `json:"credentials"`
	Host               string `json:"host"`
	GosumemoryURL      string `json:"gosumemory_url"`
	StreamCompanionURL string `json:"stream_companion_url"`
}

type AuthResponse struct {
	Error struct {
		Num int    `json:"num"`
		Str string `json:"str"`
	} `json:"error"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthError struct {
	Message string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("auth: %v", e.Message)
}

func CheckAuthResponse(message []byte) error {
	var json_data AuthResponse
	err := json.Unmarshal(message, &json_data)
	if err != nil {
		return err
	}
	if json_data.Error.Num != 0 {
		return AuthError{json_data.Error.Str}
	}
	return nil
}

// TODO: set default server Url here too
func SetDefaultValuesToConfiguration(config *Config) {
	if config.StreamCompanionURL == "" {
		config.StreamCompanionURL = "http://localhost:20727/json"
	}
	if config.GosumemoryURL == "" {
		config.GosumemoryURL = "http://localhost:24050/json"
	}
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
	SetDefaultValuesToConfiguration(&config)
	return config, err
}

func handle_command(command string, urls []string) ([]byte, error) {
	if command == "np" {
		return getOsuData(urls), nil
	}
	return nil, nil
}

func connection_handler(url string, osu_urls []string, credentials []byte) error {
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

	_, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	err = CheckAuthResponse(message)
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
		response, err := handle_command(string(message), osu_urls)

		// This currently disconnects client from server, not sure if that's desirable outcome
		if err != nil {
			error_response := &ErrorResponse{err.Error()}
			bytes_response, _ := json.Marshal(error_response)
			conn.WriteMessage(
				websocket.TextMessage,
				bytes_response,
			)
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
