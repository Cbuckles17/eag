package userinput

import (
    "os"
    "bufio"
    "strings"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// GetString will get user input and put it into a string.
func GetString() (string, error) {
    stdinReader := bufio.NewReader(os.Stdin)

    userInput, err := stdinReader.ReadString('\n')
    if err != nil {
    // checking if ReadString returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "ReadString",
        }).Errorf("Error in GetString")

        return "", err
    }

    // trim off the newline at the end of the string
    return strings.TrimSuffix(userInput, "\n"), nil
}

// GetRune will get user input and put it into a rune.
func GetRune() (rune, error) {
    stdinReader := bufio.NewReader(os.Stdin)

    userInput, _, err := stdinReader.ReadRune()
    if err != nil {
    // checking if ReadRune returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "ReadRune",
        }).Errorf("Error in GetRune")

        return 0, err
    }

    return userInput, nil
}
