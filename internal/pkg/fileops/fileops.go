package fileops

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"

    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// Create will create a file given a fully qualified file
// path and if it does not already exist.
func Create() {
    fmt.Print("FQ File Path: ")
    filePath, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
    // the file does not exist
        newFile, err := os.Create(filePath)
        if err != nil {
        // checking if os.Create returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "os.Create",
            }).Errorf("Error in Create")

            return
        }

        defer newFile.Close()

        log.WithFields(log.Fields{
            "filePath": filePath,
            "fileOp": "Create",
            "command": "os.Create",
            "arguments": filePath,
        }).Infof("File Created Successfully")
    } else {
    // file already exists
        log.WithFields(log.Fields{
            "filePath": filePath,
            "fileOp": "Create",
        }).Infof("File Already Exists")
    }
}

// Delete will delete a file given a fully qualified file
// path and if it exists.
func Delete() {
    fmt.Print("FQ File Path: ")
    filePath, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    if err := os.Remove(filePath); err != nil {
    // checking if os.Remove returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "os.Remove",
        }).Errorf("Error in Delete")

        return
    }

    log.WithFields(log.Fields{
        "filePath": filePath,
        "fileOp": "Delete",
        "command": "os.Remove",
        "arguments": filePath,
    }).Infof("File Deleted Successfully")
}

// Write will write to a file given a fully qualified file
// path and if it exists. It will also overwrite the file 
// dependant on the passed in argument.
func Write(writeType string) {
    fmt.Print("FQ File Path: ")
    filePath, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    openedFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
    // checking if os.OpenFile returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "os.OpenFile",
        }).Errorf("Error in Write")

        return 
    } 

    defer openedFile.Close()

    if strings.Compare(writeType, "Overwrite") == 0 {
    // truncate file
        if err := openedFile.Truncate(0); err != nil {
        // checking if Truncate returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "Truncate",
            }).Errorf("Error in Write")

            return 
        } 
    } else if strings.Compare(writeType, "Append") != 0 {
        log.WithFields(log.Fields{
            "err": err,
            "cause": "Invalid writeType",
        }).Errorf("Error in Write")

        return 
    }

    fmt.Print("Enter Data: ")
    dataToWrite, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    if _, err := openedFile.WriteString(dataToWrite + "\n"); err != nil {
    // checking if WriteString returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "WriteString",
        }).Errorf("Error in Write")

        return
    }

    log.WithFields(log.Fields{
        "filePath": filePath,
        "fileOp": "Write-" + writeType,
        "command": "WriteString",
        "arguments": dataToWrite,
    }).Infof("File Written Successfully")
}

// Read will read from a file given a fully qualified file
// path and if it exists.
func Read() {
    fmt.Print("FQ File Path: ")
    filePath, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    dataRead, err := ioutil.ReadFile(filePath)
    if err != nil {
    // checking if ioutil.ReadFile returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "ioutil.ReadFile",
        }).Errorf("Error in Read")

        return
    }

    log.WithFields(log.Fields{
        "filePath": filePath,
        "fileOp": "Read",
        "command": "ioutil.ReadFile",
        "arguments": filePath,
        "dataRead": string(dataRead),
    }).Infof("File Read Successfully")
}
