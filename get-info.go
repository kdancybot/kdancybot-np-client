package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "encoding/json"
    // "sync"
)

// type OsuDataResponse struct {

// }

// func getDataFromUrl(url string, wg *sync.WaitGroup, out chan<- bytes[]) ([]byte, error) {
func getDataFromUrl(url string) ([]byte, error) {
    // this is kinda bad because for gosumemory delay will be at the very least 2s
    client := http.Client{
        Timeout: time.Millisecond * 2000, // this will almost never timeout
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

    log.Print("Successfully read data from ", url)
    return data, nil
}


// TODO: Rewrite to use goroutines
func getOsuData(urls []string) ([]byte) {
    var errors []string
    for _, url := range urls {
        data, err := getDataFromUrl(url)
        if err == nil {
            return data
        } else {
            log.Print(err)
            errors = append(errors, err.Error())
        }
    }
    errors_map := map[string]interface{}{}
    errors_map["error"] = errors
    errors_json, _ := json.Marshal(errors_map)

    log.Print(string(errors_json))
    return errors_json
}

    // var wg sync.WaitGroup
    // wg.add(1)
    // bytes_chan := make(chan bytes[], len(urls))
    // for i, url := range urls {
    //     go getDataFromUrl(url, wg, bytes_chan)
    // }
    // out := <-bytes_chan
    // return out, 

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
