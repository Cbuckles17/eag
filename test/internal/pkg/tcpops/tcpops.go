package tcpops

import (
    "net"
    "time"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

func Listen(address string) {
    listener, err := net.Listen("tcp", address)
    if err != nil {
    // checking if net.Listen returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.Listen",
        }).Errorf("Error in Listen")

    }

    log.WithFields(log.Fields{
        "address": address,
    }).Info("Listener Started")

    defer listener.Close()

    //loop to continue accepting connections
    for {
        conn, err := listener.Accept()
        if err != nil {
        // checking if Accept returned an error
            log.WithFields(log.Fields{
                "err": err,
                "cause": "Accept",
            }).Warnf("Warn in Listen")

            //don't die
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // magic number of 1024
    buffer := make([]byte, 1024)
    
    dataSize, err := conn.Read(buffer)
    if err != nil {
    // checking if Read returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "Read",
        }).Errorf("Error in handleConnection")

        return
    }

    log.WithFields(log.Fields{
        "packetSize": dataSize,
        "packetData": string(buffer[:dataSize]),
    }).Info("Packet Received")

    // sleep the connection to simulate a connection that takes time
    // to send a response
    time.Sleep(5 * time.Second)

    if _, err = conn.Write([]byte("Hello TCP client!\n")); err != nil {
    // checking if Write returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "Write",
        }).Errorf("Error in handleConnection")
    }

    conn.Close()
}
