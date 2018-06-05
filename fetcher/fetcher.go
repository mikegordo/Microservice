package main

import (
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "os"
    "strconv"
)

const POST_URL = "http://localhost:8080/"
const LENGTH = 8
const NEED_CAPS = false

func main() {    
    setup()
    message, length, needCaps := buildMessage()

    log.Printf("requesting length %d, caps: %v", length, needCaps);

    value := fetchDataFromServer(message)

    log.Printf("response from generator: %s - storing to Redis", value);    

    storeToRedis(value, length, needCaps)
}

/**
* Build json message
*/
func buildMessage() ([]byte, int, bool) {
    length := LENGTH
    if os.Getenv("FETCHER_LENGTH") != "" {
        length, _ = strconv.Atoi(os.Getenv("FETCHER_LENGTH"))
    }

    needCaps := NEED_CAPS
    if os.Getenv("FETCHER_CAPS") != "" {
        needCaps, _ = strconv.ParseBool(os.Getenv("FETCHER_CAPS"))
    }
    
    messageElement := RequestFormat{&length, &needCaps}
    message, err := json.Marshal(messageElement)
    if err != nil {
        log.Printf("Error: %v", err.Error())
        panic(err.Error())
    }

    return message, length, needCaps
}

/**
* Make a POST call and process the response
*/
func fetchDataFromServer(message []byte) string {
    /* build request */
    req, err := http.NewRequest("POST", POST_URL, bytes.NewBuffer(message))
    if err != nil {
        log.Printf("Error: %v", err.Error())
        panic(err.Error())
    }

    req.Header.Set("X-Custom-Header", "himor")
    req.Header.Set("Content-Type", "application/json")

    /* submit request*/
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error: %v", err.Error())
        panic(err.Error())
    }

    defer resp.Body.Close()

    if resp.Status != "200 OK" {
        log.Printf("Response status: %v", resp.Status)
        log.Printf("Response headers: %v", resp.Header)
        panic("")
    }

    /* parse response */
    body,  _ := ioutil.ReadAll(resp.Body)
    var serverReturn Return
    err = json.Unmarshal(body, &serverReturn)
    if err != nil {
        log.Printf("Error: %v", err.Error())
        panic(err.Error())
    }

    if !serverReturn.Status {
        log.Printf("Error: code %d", serverReturn.ErrorCode)
        panic("Error: " + string(serverReturn.ErrorCode))
    }

    return serverReturn.Value
}
