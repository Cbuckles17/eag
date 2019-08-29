package binaryops

import (
    "os/exec"
    "bytes"
    "regexp"
    "fmt"

    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// Execute will run a binary/command. It can wait for the response
// dependant on the passed in argument.
func Execute(wait bool) {
    var binaryOut bytes.Buffer

    fmt.Print("Binary: ")
    binaryPath, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    fmt.Print("Arguments: ")
    arguments, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    //regex to split command args by whitespace unless they're encased in quotes
    regexSplitSpacesIgnoreQuotes := regexp.MustCompile(`"[^"]*"|[^ ]+`)
    splitArguments := regexSplitSpacesIgnoreQuotes.FindAllString(arguments, -1)

    binaryCommand := exec.Command(binaryPath, splitArguments...)

    if wait {
        binaryCommand.Stdout = &binaryOut
        binaryCommand.Stderr = &binaryOut

        if err := binaryCommand.Run(); err != nil {
        // checking if Run returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "Run",
            }).Errorf("Error in Execute-Wait")

            return
        }

        log.WithFields(log.Fields{
            "binaryOp": "Execute-Wait",
            "command": binaryPath,
            "arguments": splitArguments,
            "pid":   binaryCommand.Process.Pid,
            "binaryOut": binaryOut.String(),
        }).Infof("Executed Binary Successfully")
    } else {
        if err := binaryCommand.Start(); err != nil {
        // checking if Start returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "Start",
            }).Errorf("Error in Execute-NoWait")

            return
        }

        log.WithFields(log.Fields{
            "binaryOp": "Execute-NoWait",
            "command": binaryPath,
            "arguments": splitArguments,
            "pid":   binaryCommand.Process.Pid,
        }).Infof("Executed Binary Successfully")
    }
}