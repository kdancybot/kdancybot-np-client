package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "log"
)

func getDataFromGosumemory(url string) ([]byte, error) {
    client := http.Client{
        Timeout: time.Millisecond * 500,
    }
    response, err := client.Get(url)
    // response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    log.Print("Successfully read data from gosumemory")
    return data, nil
}

// func getData()([]byte, error) {
//     jsonData, err := getDataFromGosumemory()
//     if err != nil {
//         log.Fatal("Error getting and unmarshaling JSON data:", err)
//     }

//     // log.Printf("Received JSON data: %+v", jsonData)
//     // log.Printf(jsonData["message"].(string))
// 	// log.Printf(jsonData["status"].(string))
//     return jsonData
// }
