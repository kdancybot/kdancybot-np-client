package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "log"
    // "sync"
)

// type OsuDataResponse struct {

// }

// func getDataFromUrl(url string, wg *sync.WaitGroup, out chan<- bytes[]) ([]byte, error) {
func getDataFromUrl(url string) ([]byte, error) {
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

    log.Print("Successfully read data from ", url)
    return data, nil
}

func getOsuData(urls []string) ([]byte, error) {
    var err error
    for _, url := range urls {
        data, err := getDataFromUrl(url)
        if err == nil {
            return data, nil
        }
        else {
            log.Print(err)
        }
    }
    return nil, err

    // var wg sync.WaitGroup
    // wg.add(1)
    // bytes_chan := make(chan bytes[], len(urls))
    // for i, url := range urls {
    //     go getDataFromUrl(url, wg, bytes_chan)
    // }
    // out := <-bytes_chan
    // return out, 
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
