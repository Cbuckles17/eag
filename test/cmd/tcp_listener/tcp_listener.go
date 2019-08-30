package main

import (
    "os"
    "fmt"
    stdlog "log"

    "eag/test/internal/pkg/tcpops"
    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// init is used to set up logging.
func init() {
    config := log.Conf{
        LogFormat:          "JSON",
        LogLevel:           log.DEBUGLEVEL,
        LogTimeUTC:         true,
        LogToStdout:        true,
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
