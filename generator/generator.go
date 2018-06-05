package main

import (
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    setup()

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t RequestFormat

    w.Header().Set("Content-Type", "application/json")

    if (nil != decoder.Decode(&t)) {
        returnError(w, 400, "Incorrect Json")
        return
    }

    length := 8
    if (t.Length != nil) {
        length = *t.Length
    }

    includeCaps := true
    if (t.IncludeCaps != nil) {
        includeCaps = *t.IncludeCaps
    }

    output := generateRandomString(length, includeCaps)

    returnJson(w, Return{true, 200, output})
}

/**
* Returns error response
*/
func returnError(w http.ResponseWriter, code int, err string) {
    log.Printf("Error: %d %s", code, err)
    errorReturn := Return{false, code, err}
    js, e := json.Marshal(errorReturn)
    if e != nil {
        log.Printf("Error: %v", e.Error())
    }
    w.Write(js)
}

/**
* Return correct response
*/
func returnJson(w http.ResponseWriter, element Return) {
    js, err := json.Marshal(element)
    if err != nil {
        log.Printf("Error: %v", err.Error())
    }
    w.Write(js)
}