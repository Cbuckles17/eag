package main

import (
    "os"
    "flag"
    "fmt"
    stdlog "log"

    "eag/test/internal/pkg/tcpops"
    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// init is used to set up logging.
func init() {
    logToFile := false

    // set up flag for enabling logging to file
    logFilePath := flag.String("l", "", "Path to a log file. It will be created if it DNE.")
    flag.Parse()

    if len(*logFilePath) > 0 {
    // set to log a file
        logToFile = true
    }

    config := log.Conf{
        LogFormat:          "JSON",
        LogLevel:           log.DEBUGLEVEL,
        LogTimeUTC:         true,
        LogToStdout:        true,
        LogToFile:          logToFile,
        FilePath:           *logFilePath,
    }
    
    if _, err := log.Create(config); err != nil {
        stdlog.Fatalf("Failed Creating Log: %s\n", err)
    }
}

// main starts the listener.
func main() {
    fmt.Print("Address: ")
    address, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        os.Exit(-1)
    }

    //start the listener
    tcpops.Listen(address)
}
