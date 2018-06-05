package main

import (
    "log"
    "io"
    "os"
)

const LOGFILE        = "fetcher.log"
const LOG_TO_FILE    = 1
const LOG_TO_CONSOLE = 2

/**
* Set up everything
*/
func setup() {
    setupLogging(LOG_TO_CONSOLE | LOG_TO_FILE)
}

/**
* Set up logging
*/
func setupLogging(source int) {
    if source == LOG_TO_FILE | LOG_TO_CONSOLE {
        f, err := os.OpenFile(LOGFILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        if err != nil {
            log.Fatalf("Error opening logfile: %v", err)
        }
        multi := io.MultiWriter(f, os.Stdout)
        log.SetOutput(multi)
    } else if source == LOG_TO_FILE {
        f, err := os.OpenFile(LOGFILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        if err != nil {
            log.Fatalf("Error opening logfile: %v", err)
        }
        log.SetOutput(f)
    }
}
