package udpops

import (
    "net"
    "time"

    log "github.com/Cbuckles17/genericlog/pkg/genericlog"
)

func Listen(address string) {
    resolvedAddress, err := net.ResolveUDPAddr("udp", address)
    if err != nil {
    // checking if net.ResolveUDPAddr returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.ResolveUDPAddr",
        }).Errorf("Error in Listen")

        return
    }

    connections, err := net.ListenUDP("udp", resolvedAddress)
    if err != nil {
    // checking if net.Listen returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "net.ListenUDP",
        }).Errorf("Error in Listen")

    }

    log.WithFields(log.Fields{
        "address": address,
    }).Info("Listener Started")

    defer connections.Close()

    //loop to continue accepting connections
    for {
         // wait for UDP client to connect
         handleConnection(connections)
    }
}

func handleConnection(conn *net.UDPConn) {
    // magic number of 1024
    buffer := make([]byte, 1024)

    dataSize, sourceAddress, err := conn.ReadFromUDP(buffer)
    if err != nil {
    // checking if ReadFromUDP returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "ReadFromUDP",
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

    if _, err = conn.WriteToUDP([]byte("Hello UDP client!\n"), sourceAddress); err != nil {
    // checking if WriteToUDP returned an error
        log.WithFields(log.Fields{
            "err": err,
            "cause": "WriteToUDP",
        }).Errorf("Error in handleConnection")

        return
    }
}
