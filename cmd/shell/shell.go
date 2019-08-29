package main

import (
    "fmt"
    "os"
    "os/user"
    stdlog "log"

    "eag/internal/pkg/binaryops"
    "eag/internal/pkg/fileops"
    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// init is used to set up logging.
func init() {
    user, err := user.Current()
    if err != nil {
        stdlog.Fatalf("Failed Getting User Information: %s\n", err)
    }

    // TO DO: Pass in conf via commandline args else use defaulted values
    config := log.Conf{
        LogFormat:          "JSON",
        LogLevel:           log.DEBUGLEVEL,
        LogTimeUTC:         true,
        LogToStdout:        true,
        LogToFile:          true,
        FilePath:           "/Users/cbuckley/go/test.log",
        UseDefaultFields:   true,
        DefaultFields:      log.Fields{"user_name": user.Username, "user_id": user.Uid, "pid": os.Getpid()},
    }
    
    if _, err = log.Create(config); err != nil {
        stdlog.Fatalf("Failed Creating Log: %s\n", err)
    }
}

// main runs a loop of getting user input and deciding what to do
// from there.
func main() {
    for {
        fmt.Print("shell>")
        choice, err := userinput.GetRune()
        if err != nil {
        // checking if userinput.GetRune returned an error
            os.Exit(-1)
        }

        switch choice {
            case 'c', 'C':
                fileops.Create()
            case 'd', 'D':
                fileops.Delete()
            case 'a', 'A':
                fileops.Write("Append")
            case 'o', 'O':
                fileops.Write("Overwrite")
            case 'r', 'R':
                fileops.Read()
            case 'b', 'B':
                binaryops.Execute(false)
            case 'w', 'W':
                binaryops.Execute(true)
            // case 'h', 'H':
            //  help()
            case 'q', 'Q':
                os.Exit(0)
            default:
                fmt.Println("Invalid Command: ", string(choice))
        }
    }
}
