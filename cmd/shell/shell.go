package main

import (
    "flag"
    "fmt"
    "os"
    "os/user"
    "strings"
    stdlog "log"

    "eag/internal/pkg/binaryops"
    "eag/internal/pkg/fileops"
    "eag/internal/pkg/netops"
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

    user, err := user.Current()
    if err != nil {
        stdlog.Fatalf("Failed Getting User Information: %s\n", err)
    }

    config := log.Conf{
        LogFormat:          "JSON",
        LogLevel:           log.DEBUGLEVEL,
        LogTimeUTC:         true,
        LogToStdout:        true,
        LogToFile:          logToFile,
        FilePath:           *logFilePath,
        UseDefaultFields:   true,
        DefaultFields:      log.Fields{"user_name": user.Username, "user_id": user.Uid, "pid": os.Getpid()},
    }
    
    if _, err = log.Create(config); err != nil {
        stdlog.Fatalf("Failed Creating Log: %s\n", err)
    }
}

// help just prints all of the options out
func help() {
    fmt.Printf(
        "c : Create File\n" +
        "d : Delete File\n" +
        "a : Append File\n" +
        "o : Overwrite File\n" +
        "r : Read File\n" +
        "b : Execute Binary No Wait\n" +
        "bw: Execute Binary Wait\n" +
        "t : Write TCP No Wait\n" +
        "tw: Write TCP Wait\n" +
        "u : Write UDP No Wait\n" +
        "uw: Write UDP Wait\n")
}

// main runs a loop of getting user input and deciding what to do
// from there.
func main() {
    for {
        fmt.Print("shell>")
        choice, err := userinput.GetString()
        if err != nil {
        // checking if userinput.GetString returned an error
            os.Exit(-1)
        }

        choice = strings.ToLower(choice)

        switch choice {
            case "c":
                fileops.Create()
            case "d":
                fileops.Delete()
            case "a":
                fileops.Write("Append")
            case "o":
                fileops.Write("Overwrite")
            case "r":
                fileops.Read()
            case "b":
                binaryops.Execute(false)
            case "bw":
                binaryops.Execute(true)
            case "t":
                netops.WriteTCP(false)
            case "tw":
                netops.WriteTCP(true)
            case "u":
                netops.WriteUDP(false)
            case "uw":
                netops.WriteUDP(true)
            case "h":
                help()
            case "q":
                os.Exit(0)
            default:
                fmt.Println("Invalid Command: ", choice)
        }
    }
}
