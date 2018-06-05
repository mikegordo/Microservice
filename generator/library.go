package main

import (
    "log"
    "strings"
    "math/rand"
    "time"
)

/**
* Generate random string of desired size
* @param size int
* @param includeCaps bool
* @return string
*/
func generateRandomString(size int, includeCaps bool) string {
    if size < 1 {
        log.Fatal("String length cannot be less than 1")
    }

    var str []string
    var dictionary string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

    if !includeCaps {
        dictionary = dictionary[0:len(dictionary) / 2]
    } 

    randGenSource := rand.NewSource(time.Now().UnixNano())
    randGen := rand.New(randGenSource)
    
    for i := 0; i < size; i++ {
        str = append(str, string(dictionary[randGen.Intn(len(dictionary))])) 
    }

    return strings.Join(str, "")
}
