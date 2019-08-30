package netops

import (
    "net"
    "fmt"
    "bufio"

    "eag/internal/pkg/userinput"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

// WriteTCP will create a tcp connection and write to it. It can
// wait for the response dependant on the passed in argument.
func WriteTCP(wait bool) {
    fmt.Print("Address: ")
    address, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    conn, err := net.Dial("tcp", address)
    if err != nil {
     // checking if net.Dial returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.Dial",
        }).Errorf("Error in WriteTCP")

        return
    }

    defer conn.Close()

    fmt.Print("Data: ")
    dataToSend, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

   bytesSent, err := fmt.Fprintf(conn, dataToSend)
   if err != nil {
    // checking if Fprintf returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "Fprintf",
        }).Errorf("Error in WriteTCP")

        return
    }

    if wait {
    // wait for response after sending packet
        responseData, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
        // checking if ReadString returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "ReadString",
            }).Errorf("Error in WriteTCP")

            return
        }

        log.WithFields(log.Fields{
            "netOp": "WriteTCP-Wait",
            "protocol": "tcp",
            "source": conn.LocalAddr().String(),
            "destination:": conn.RemoteAddr().String(),
            "bytesSent": bytesSent,
            "dataSent": dataToSend,
            "responseData": responseData,
            "command": "Fprintf",
            "arguments": dataToSend,
        }).Infof("Sent TCP packet Successfully")

    } else {
    // just send packet
        log.WithFields(log.Fields{
            "netOp": "WriteTCP-NoWait", 
            "protocol": "tcp",
            "source": conn.LocalAddr().String(),
            "destination:": conn.RemoteAddr().String(),
            "bytesSent": bytesSent,
            "dataSent": dataToSend,
            "command": "Fprintf",
            "arguments": dataToSend,
        }).Infof("Sent TCP packet Successfully")
    }
}

// WriteUDP will create a udp connection and write to it. It can
// wait for the response dependant on the passed in argument.
func WriteUDP(wait bool) {
    fmt.Print("Address: ")
    address, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    resolvedAddress, err := net.ResolveUDPAddr("udp", address)
    if err != nil {
    // checking if net.ResolveUDPAddr returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.ResolveUDPAddr",
        }).Errorf("Error in WriteUDP")

        return
    }


    conn, err := net.DialUDP("udp", nil, resolvedAddress)
    if err != nil {
    // checking if net.DialUDP returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.DialUDP",
        }).Errorf("Error in WriteUDP")

        return
    }

    defer conn.Close()

    fmt.Print("Data: ")
    dataToSend, err := userinput.GetString()
    if err != nil {
    // checking if userinput.GetString returned an error
        return
    }

    bytesSent, err := fmt.Fprintf(conn, dataToSend)
    if err != nil {
    // checking if Fprintf returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "Fprintf",
        }).Errorf("Error in WriteUDP")

        return
    }

    if wait {
    // wait for response after sending packet
        // magic number of 1024
        buffer := make([]byte, 1024)

        responseSize, _, err := conn.ReadFromUDP(buffer)
        if err != nil {
        // checking if ReadFromUDP returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "ReadFromUDP",
            }).Errorf("Error in WriteUDP")

            return
        }

        log.WithFields(log.Fields{
            "netOp": "WriteUDP-Wait",
            "protocol": "udp",
            "source": conn.LocalAddr().String(),
            "destination:": conn.RemoteAddr().String(),
            "bytesSent": bytesSent,
            "dataSent": dataToSend,
            "responseData": string(buffer[:responseSize]),
            "command": "Fprintf",
            "arguments": dataToSend,
        }).Infof("Sent UDP packet Successfully")

    } else {
    // just send packet
        log.WithFields(log.Fields{
            "netOp": "WriteUDP-NoWait", 
            "protocol": "udp",
            "source": conn.LocalAddr().String(),
            "destination:": conn.RemoteAddr().String(),
            "bytesSent": bytesSent,
            "dataSent": dataToSend,
            "command": "Fprintf",
            "arguments": dataToSend,
        }).Infof("Sent UDP packet Successfully")
    }
}
