package main

import (
    "gopkg.in/redis.v4"
    "log"
    "time"
    "encoding/json"
)

const REDIS_KEY = "thekey"

func storeToRedis(str string, length int, includeCaps bool) {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    element := RedisElement {length, includeCaps, str, int32(time.Now().Unix())}
    serialized, err := json.Marshal(element)

    if err != nil {
        log.Printf("Error: %v", err.Error())
        panic(err.Error())
    }

    client.SAdd(REDIS_KEY, string(serialized))
}