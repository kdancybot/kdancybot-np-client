package main

import (
    "net/http"
    "io/ioutil"
    // "log"
    // "encoding/json"
)

func getDataFromGosumemory() ([]byte, error) {
    url := "http://127.0.0.1:24050/json"

    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    // var jsonData map[string]interface{}

    // if err := json.Unmarshal(data, &jsonData); err != nil {
    //     return nil, err
    // }

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
